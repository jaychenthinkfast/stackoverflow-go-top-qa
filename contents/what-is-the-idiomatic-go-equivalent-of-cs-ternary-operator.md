# Go 语言中与 C 语言的三元运算符等价的惯用表达方式是什么？

## 问题

在 C/C++（以及许多类似的语言）中，根据条件声明和初始化变量的常见惯用方法是使用三元条件运算符：

```
int index = val > 0 ? val : -val
```

Go 没有条件运算符。实现上述代码的最佳惯用方法是什么？我找到了以下解决方案，但它似乎相当冗长

```go
var index int

if val > 0 {
    index = val
} else {
    index = -val
}
```

有没有更好的方法？

## 回答1


如所指出的（并且希望不会令人惊讶），在 Go 中使用 if+else 确实是进行条件判断的习惯用法。

除了完整的 var + if + else 代码块之外，这种拼写也经常被使用：

```go
index := val
if val <= 0 {
    index = -val
}
```

如果您的代码块足够重复，例如相当于

> `int value = a <= b ? a : b`

然后，您可以创建一个函数来保存它：

```go
func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

...

value := min(a, b)
```

编译器将内联这些简单的函数，因此它更快、更清晰、更短。

## stackoverflow链接

https://stackoverflow.com/questions/19979178/what-is-the-idiomatic-go-equivalent-of-cs-ternary-operator
