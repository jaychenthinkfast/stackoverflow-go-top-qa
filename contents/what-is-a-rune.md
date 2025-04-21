# 什么是符文？(rune)

## 问题


Go 中的 rune 是什么？

我一直在谷歌搜索，但 Golang 只说了一句： rune 是 int32 的别名。

但为什么整数被用在像交换大小写这样的地方呢？

以下是一个 swapcase 函数。所有的 <= 和 - 是什么意思？

那为什么 switch 没有任何参数呢？

&& 应该意味着和，但 r <= 'z' 是什么意思？

```go
func SwapRune(r rune) rune {
    switch {
    case 'a' <= r && r <= 'z':
        return r - 'a' + 'A'
    case 'A' <= r && r <= 'Z':
        return r - 'A' + 'a'
    default:
        return r
    }
}
```

```go
func SwapCase(str string) string {
    return strings.Map(SwapRune, str)
}
```

我明白这是将 rune 映射到 string 以返回交换后的字符串。但我不明白 rune 或 byte 在这里是如何工作的。

## 回答1



符文字面量只是 32 位整数值（然而它们是无类型的常量，所以它们的类型可以改变）。它们代表 Unicode 码点。例如，符文字面量 'a' 实际上是数字 97 。

因此，你的程序基本上等同于：

```go
package main

import "fmt"

func SwapRune(r rune) rune {
    switch {
    case 97 <= r && r <= 122:
        return r - 32
    case 65 <= r && r <= 90:
        return r + 32
    default:
        return r
    }
}

func main() {
    fmt.Println(SwapRune('a'))
}
```

应该很明显，如果你查看 Unicode 映射，该范围与 ASCII 相同。此外，32 实际上是字符的大写和小写码点之间的偏移量。所以通过给 'A' 加上 32 ，你得到 'a' ，反之亦然。

## stackoverflow链接

https://stackoverflow.com/questions/19310700/what-is-a-rune
