# Go中如何高效连接字符串

## 问题
在Go中，string是基本类型，这意味着它是只读的，并且每个对它的操作将会创建一个新的string

如果我想多次连接字符串但是不知道生成字符串的长度，有什么最佳实践?

简单方案：
``` 
var s string
for i := 0; i < 1000; i++ {
    s += getShortStringFromSomewhere()
}
return s
```
但是看起来并不高效。

## 回答1
```
package main

import (
    "bytes"
    "strings"
    "testing"
)

const (
    sss = "xfoasneobfasieongasbg"
    cnt = 10000
)

var (
    bbb      = []byte(sss)
    expected = strings.Repeat(sss, cnt)
)

func BenchmarkCopyPreAllocate(b *testing.B) {
    var result string
    for n := 0; n < b.N; n++ {
        bs := make([]byte, cnt*len(sss))
        bl := 0
        for i := 0; i < cnt; i++ {
            bl += copy(bs[bl:], sss)
        }
        result = string(bs)
    }
    b.StopTimer()
    if result != expected {
        b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
    }
}

func BenchmarkAppendPreAllocate(b *testing.B) {
    var result string
    for n := 0; n < b.N; n++ {
        data := make([]byte, 0, cnt*len(sss))
        for i := 0; i < cnt; i++ {
            data = append(data, sss...)
        }
        result = string(data)
    }
    b.StopTimer()
    if result != expected {
        b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
    }
}

func BenchmarkBufferPreAllocate(b *testing.B) {
    var result string
    for n := 0; n < b.N; n++ {
        buf := bytes.NewBuffer(make([]byte, 0, cnt*len(sss)))
        for i := 0; i < cnt; i++ {
            buf.WriteString(sss)
        }
        result = buf.String()
    }
    b.StopTimer()
    if result != expected {
        b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
    }
}

func BenchmarkCopy(b *testing.B) {
    var result string
    for n := 0; n < b.N; n++ {
        data := make([]byte, 0, 64) // same size as bootstrap array of bytes.Buffer
        for i := 0; i < cnt; i++ {
            off := len(data)
            if off+len(sss) > cap(data) {
                temp := make([]byte, 2*cap(data)+len(sss))
                copy(temp, data)
                data = temp
            }
            data = data[0 : off+len(sss)]
            copy(data[off:], sss)
        }
        result = string(data)
    }
    b.StopTimer()
    if result != expected {
        b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
    }
}

func BenchmarkAppend(b *testing.B) {
    var result string
    for n := 0; n < b.N; n++ {
        data := make([]byte, 0, 64)
        for i := 0; i < cnt; i++ {
            data = append(data, sss...)
        }
        result = string(data)
    }
    b.StopTimer()
    if result != expected {
        b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
    }
}

func BenchmarkBufferWrite(b *testing.B) {
    var result string
    for n := 0; n < b.N; n++ {
        var buf bytes.Buffer
        for i := 0; i < cnt; i++ {
            buf.Write(bbb)
        }
        result = buf.String()
    }
    b.StopTimer()
    if result != expected {
        b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
    }
}

func BenchmarkBufferWriteString(b *testing.B) {
    var result string
    for n := 0; n < b.N; n++ {
        var buf bytes.Buffer
        for i := 0; i < cnt; i++ {
            buf.WriteString(sss)
        }
        result = buf.String()
    }
    b.StopTimer()
    if result != expected {
        b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
    }
}

func BenchmarkConcat(b *testing.B) {
    var result string
    for n := 0; n < b.N; n++ {
        var str string
        for i := 0; i < cnt; i++ {
            str += sss
        }
        result = str
    }
    b.StopTimer()
    if result != expected {
        b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
    }
}
``` 
执行环境: OS X 10.11.6, 2.2 GHz Intel Core i7

测试结果：
```
BenchmarkCopyPreAllocate-8         20000             84208 ns/op          425984 B/op          2 allocs/op
BenchmarkAppendPreAllocate-8       10000            102859 ns/op          425984 B/op          2 allocs/op
BenchmarkBufferPreAllocate-8       10000            166407 ns/op          426096 B/op          3 allocs/op
BenchmarkCopy-8                    10000            160923 ns/op          933152 B/op         13 allocs/op
BenchmarkAppend-8                  10000            175508 ns/op         1332096 B/op         24 allocs/op
BenchmarkBufferWrite-8             10000            239886 ns/op          933266 B/op         14 allocs/op
BenchmarkBufferWriteString-8       10000            236432 ns/op          933266 B/op         14 allocs/op
BenchmarkConcat-8                     10         105603419 ns/op        1086685168 B/op    10000 allocs/op 
```
结论:

1. CopyPreAllocate最快；AppendPreAllocate次之，但是更易编码。
2. Concat不管是运行速度还是内存使用都表现最差。
3. Buffer#Write 和 Buffer#WriteString 表现基本相当。
4. bytes.Buffer 基本使用了 Copy相同的方案。
5. Copy 和 Append 使用了和 bytes.Buffer一样的 初始长度64。
6. Append 使用了更多的内存分配，和它的扩容算法有关。

建议：

1. 对于简单场景，推荐使用Append 或 AppendPreAllocate，速度快、编码简单。
2. 对于同时需要读写buffer的场景，当然推荐使用bytes.Buffer。

## stackoverflow链接
https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go

## 扩展
在题主问题提出后Go1.10+ 增加了strings.Builder，在bytes源码注释中标明：strings.Builder比bytes.Buffer构建字符串上更加高效。
```
// String returns the contents of the unread portion of the buffer
// as a string. If the Buffer is a nil pointer, it returns "<nil>".
//
// To build strings more efficiently, see the strings.Builder type.
func (b *Buffer) String() string {
	if b == nil {
		// Special case, useful in debugging.
		return "<nil>"
	}
	return string(b.buf[b.off:])
}
```