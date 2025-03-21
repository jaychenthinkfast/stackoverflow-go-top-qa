# 格式化Go字符串而不打印？

## 问题

有没有一种简单的方法来格式化Go中的字符串而不打印字符串？

我可以做:

```go
bar := "bar"
fmt.Printf("foo: %s", bar)
```

但我想格式化的字符串返回，而不是打印，所以我可以进一步操纵它。

我也可以做一些类似的事情:

```go
s := "foo: " + bar
```

但是，当格式字符串很复杂时，这就变得很难阅读，而当一个或多个部分不是字符串并且必须先转换时，这就很麻烦了，例如

```go
i := 25
s := "foo: " + strconv.Itoa(i)
```

有没有更简单的方法来做到这一点？

## 回答1

[Sprintf](https://golang.org/pkg/fmt/#Sprintf)是你正在寻找的。

**示例**

```go
fmt.Sprintf("foo: %s", bar)
```

您还可以在使用中看到它[错误示例](https://tour.golang.org/methods/19)作为 “Go之旅” 的一部分。

```go
return fmt.Sprintf("at %v, %s", e.When, e.What)
```

## stackoverflow链接

https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing
