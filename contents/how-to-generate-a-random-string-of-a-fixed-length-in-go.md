# 如何在 Go 中生成一个固定长度的随机字符串？

## 问题

我想得到一个只包含字符（大写或小写）的随机字符串，不包含数字，在 Go 中，最快和最简单的方法是什么？

## 回答1
### 方法 1
这是一个简单通用的解决方案，使用 runes
```go
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
```
### 方法 2
使用字节，因为组成随机字符串的字符仅包含英文字母的大小写，只需使用字节即可，
因为英文字母在 UTF-8 编码中映射为字节 1-to-1（这是 Go 存储字符串的方式）。
```go

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
b := make([]byte, n)
for i := range b {
b[i] = letterBytes[rand.Intn(len(letterBytes))]
}
return string(b)
}
```
现在已经是一个很大的改进了：我们可以将其实现为 const （有 string 个常量，但没有切片常量）。
作为额外的收益，表达式 len(letters) 也将是 const !（如果 s 是字符串常量，则表达式 len(s) 是常量。）
### 方法 3
之前的解决方案通过调用 rand.Intn() 来获取随机数，以指定一个随机字母，该调用委托给 Rand.Intn() ，然后委托给 Rand.Int31n() 。

与产生 63 个随机位的随机数的 rand.Int63() 相比，这要慢得多。

所以我们只需调用 rand.Int63() 并使用除以 len(letterBytes) 后的余数即可：

这工作得很好，速度显著更快，缺点是所有字母的概率不会完全相同（假设 rand.Int63() 以相同的概率产生所有 63 位数字）。尽管失真非常小，因为字母 52 的数量远远小于 1<<63 - 1 ，所以在实际应用中这是完全可以接受的。

```go
func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
```
### 方法 4 
使用位操作，在前一个解决方案的基础上，我们可以通过仅使用表示字母数量所需的最少最低位随机数来保持字母的均匀分布。例如，如果我们有 52 个字母，则需要 6 位来表示它： 52 = 110100b 。因此，我们只使用 rand.Int63() 返回的数字的最低 6 位。为了保持字母的均匀分布，
我们只“接受”落在 0..len(letterBytes)-1 范围内的数字。如果最低位大于，我们将其丢弃并查询新的随机数。

注意，一般情况下，最低位大于或等于 len(letterBytes) 的概率小于 0.5 （平均为 0.25 ），这意味着即使这种情况发生，重复这种“罕见”情况也会降低找不到好数字的概率。经过 n 次重复后，我们仍然没有好索引的概率远小于 pow(0.5, n) ，这只是一个上限估计。对于 52 个字母，6 个最低位不是好的概率仅为 (64-52)/64 = 0.19 ；这意味着例如，在重复 10 次后没有好数字的概率为 1e-8 。

```go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
    b := make([]byte, n)
    for i := 0; i < n; {
        if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i++
        }
    }
    return string(b)
}
```
### 方法 5
之前的解决方案只使用了由 rand.Int63() 返回的 63 个随机比特中的最低 6 位。这是浪费，因为获取随机比特是我们算法中最慢的部分。

如果我们有 52 个字母，那么意味着 6 位编码一个字母索引。所以 63 个随机比特可以指定 63/6 = 10 个不同的字母索引。让我们使用这 10 个：
```go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImpr(n int) string {
    b := make([]byte, n)
    // A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
    for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = rand.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}
```
### 方法 6
前面方案已经没有太多改进的地方，可以改进但是会显著增加复杂度，不是太值得。

现在让我们找些其他可以改进的地方。随机数的来源。

存在一个提供 Read(b []byte) 功能的 crypto/rand 包，因此我们可以使用它通过单个调用获取我们需要的尽可能多的字节。这在性能方面没有帮助，因为 crypto/rand 实现了一个密码学安全的伪随机数生成器，所以它要慢得多。

因此让我们坚持使用 math/rand 包。 rand.Rand 使用 rand.Source 作为随机位的来源。 rand.Source 是一个指定了 Int63() int64 方法的接口：正是我们最新解决方案中需要和使用的唯一东西。

因此我们实际上并不需要一个 rand.Rand （无论是显式的还是 rand 包的全局、共享的那个），对我们来说，一个 rand.Source 就足够了：
```go
var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}
```
请注意，此最后解决方案不需要您初始化（种子） math/rand 包的全局 Rand ，因为该功能未使用（并且我们的 rand.Source 已正确初始化/种子）。

此处还需注意一点： math/rand 的 package doc 中声明：默认源对多个 goroutine 的并发使用是安全的。

因此，默认源比通过 Source 获得的源慢，因为默认源必须在并发访问/使用下提供安全性，而 rand.NewSource() 不提供这种（因此它返回的 Source 可能更快）。

### 方法 7
使用strings.Builder，Go 1.10 引入了 strings.Builder 。 strings.Builder 是一种我们可以用来构建类似于 bytes.Buffer 的 string 内容的新类型。
内部它使用 []byte 来构建内容，完成后，我们可以使用它的 Builder.String() 方法获取最终的 string 值。但其中酷的地方在于，它执行了上述我们刚刚提到的复制操作，而不进行复制。
它敢于这样做，因为用于构建字符串内容的字节切片没有被暴露，因此可以保证没有人会无意或恶意地修改它，以改变产生的“不可变”字符串。

所以我们的下一个想法是不在切片中构建随机字符串，而是借助 strings.Builder ，这样一旦完成，我们就可以获取并返回结果，而无需复制它。
这可能在速度方面有所帮助，并且肯定有助于内存使用和分配。
```go
func RandStringBytesMaskImprSrcSB(n int) string {
    sb := strings.Builder{}
    sb.Grow(n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            sb.WriteByte(letterBytes[idx])
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return sb.String()
}
```
请注意，在创建新的 strings.Buidler 之后，我们调用了它的 Builder.Grow() 方法，确保它分配足够大的内部切片（以避免在添加随机字母时进行重新分配）。

### 方法 8
strings.Builder 在内部构建字符串，与我们自己做的相同。所以基本上通过 strings.Builder 做会有一些开销，我们切换到 strings.Builder 的唯一目的是为了避免切片的最终复制。

strings.Builder 通过使用包 unsafe 避免了最终的复制：

我们也可以自己这样做。所以这里的想法是切换回在 []byte 中构建随机字符串，但当我们完成时，不要将其转换为 string 返回，而进行一个不安全的转换：获取一个 string ，它指向我们的字节切片作为字符串数据。
### 完整代码
```go
func RandStringBytesMaskImprSrcUnsafe(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return *(*string)(unsafe.Pointer(&b))
}
```



完整代码：
```go
package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"
	"unsafe"
)

// Implementations

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// Benchmark functions

const n = 16

func BenchmarkRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringRunes(n)
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytes(n)
	}
}

func BenchmarkBytesRmndr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesRmndr(n)
	}
}

func BenchmarkBytesMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMask(n)
	}
}

func BenchmarkBytesMaskImpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpr(n)
	}
}

func BenchmarkBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrc(n)
	}
}
func BenchmarkBytesMaskImprSrcSB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrcSB(n)
	}
}

func BenchmarkBytesMaskImprSrcUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrcUnsafe(n)
	}
}

```
### 基准测试结果
```go
BenchmarkRunes-4                     2000000    723 ns/op   96 B/op   2 allocs/op
BenchmarkBytes-4                     3000000    550 ns/op   32 B/op   2 allocs/op
BenchmarkBytesRmndr-4                3000000    438 ns/op   32 B/op   2 allocs/op
BenchmarkBytesMask-4                 3000000    534 ns/op   32 B/op   2 allocs/op
BenchmarkBytesMaskImpr-4            10000000    176 ns/op   32 B/op   2 allocs/op
BenchmarkBytesMaskImprSrc-4         10000000    139 ns/op   32 B/op   2 allocs/op
BenchmarkBytesMaskImprSrcSB-4       10000000    134 ns/op   16 B/op   1 allocs/op
BenchmarkBytesMaskImprSrcUnsafe-4   10000000    115 ns/op   16 B/op   1 allocs/op
```
仅通过将rune切换为字节，我们立即获得 24%的性能提升，内存需求降低到三分之一。

rand.Int63() 代替 rand.Intn() 可以再提升 20%。

位操作（以及在索引较大时重复）会稍微减慢速度（由于重复调用）：-22%...

但当我们利用所有（或大部分）的 63 个随机位（来自一个 rand.Int63() 调用的 10 个索引）：速度大幅提升：3 倍。

如果我们使用（非默认，新的） rand.Source 而不是 rand.Rand ，我们又能再获得 21%。

如果我们使用 strings.Builder ，我们可以在速度上获得微小的 3.5%提升，同时我们还实现了 50%的内存使用和分配减少！这很不错！

如果我们敢于使用包 unsafe 而不是 strings.Builder ，我们又能获得 14%的提升。

将最终方案与初始方案进行比较： RandStringBytesMaskImprSrcUnsafe() 比 RandStringRunes() 快 6.3 倍，内存使用减少六分之一，分配次数减少一半。

## stackoverflow链接

https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
