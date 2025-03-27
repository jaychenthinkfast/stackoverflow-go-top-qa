# Go 中测试空字符串的最佳方法是什么？

## 问题

哪种方法（最符合 Go 风格）用于测试非空字符串？

```go
if len(mystring) > 0 { }
```

Or:

```go
if mystring != "" { }
```

或者其他方法？

## 回答1


这两种风格都在 Go 标准库中使用。

```go
if len(s) > 0 { ... }
```

可在 strconv 包中找到：http://golang.org/src/pkg/strconv/atoi.go

```go
if s != "" { ... }
```

可在 encoding/json 包中找到：http://golang.org/src/pkg/encoding/json/encode.go

两者都很地道且足够清晰。这更多是个人品味和清晰度的问题。

Russ Cox 在[ golang-nuts](https://groups.google.com/forum/#!topic/golang-nuts/7Ks1iq2s7FA) 线程中写道：

> 让代码更清晰的那个
> 如果我即将查看元素 x，我通常会写
> 即使 x == 0，但如果我关心
> 这是否是特定的字符串”我倾向于写 s == ""。
>
> 有理由假设成熟的编译器会将
> len(s) == 0 且 s == "" 编译成相同的、高效的代码。
> ...
>
> 让代码清晰易懂。

如 [Timmmm](https://stackoverflow.com/a/54804858/694331) 的回答中指出的，Go 编译器在这两种情况下确实会生成相同的代码。


## stackoverflow链接

https://stackoverflow.com/questions/18594330/what-is-the-best-way-to-test-for-an-empty-string-in-go
