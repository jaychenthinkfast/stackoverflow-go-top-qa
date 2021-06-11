# Go中什么是表示枚举的惯用方式

## 问题
我正在尝试表示一个简化的染色体，它由N个基础组成，每个都是{A, C, T, G}中的一个。

我想通过枚举的方式进行限定，但是我想知道Go中最惯用的枚举方式是什么？

## 回答1
引用自[Go语言规范](https://golang.org/ref/spec#Iota)
    
    预定义标识符iota表示连续的无类型整数常量。它的值根据常量表达式变化，从0开始。常用于构建一组相关的常量集。
```
const (  // iota is reset to 0
        c0 = iota  // c0 == 0
        c1 = iota  // c1 == 1
        c2 = iota  // c2 == 2
)

const (
        a = 1 << iota  // a == 1 (iota has been reset)
        b = 1 << iota  // b == 2
        c = 1 << iota  // c == 4
)

const (
        u         = iota * 42  // u == 0     (untyped integer constant)
        v float64 = iota * 42  // v == 42.0  (float64 constant)
        w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0 (iota has been reset)
const y = iota  // y == 0 (iota has been reset)
```
    在一个表达式列表中，每个iota值是一样的。
``` 
const (
        bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
        bit1, mask1                           // bit1 == 2, mask1 == 1
        _, _                                  // skips iota == 2
        bit3, mask3                           // bit3 == 8, mask3 == 7
)
```
所以你可以这样编码：
``` 
const (
        A = iota
        C
        T
        G
)
```
或者
``` 
type Base int

const (
        A Base = iota
        C
        T
        G
)
```
如果你希望使用基于int的单独类型。
## stackoverflow链接
https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go