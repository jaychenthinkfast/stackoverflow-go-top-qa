# 如何避免令人烦恼的错误“已声明但未使用”

## 问题

我在学习 Go 语言，但感觉有点烦恼，编译时不能留下任何未使用的变量或包。

这真的让我效率降低。例如，我只是想声明一个新的包，打算稍后使用，或者取消注释一些命令来测试。我总是遇到这个错误，需要去取消所有这些注释。

在 Go 语言中有没有办法避免这种检查？

## 回答1

这个错误是为了迫使你写出更好的代码，并确保使用你声明的或导入的每一项。这使得阅读他人编写的代码更容易（你总是确信所有声明的变量都将被使用），并避免一些可能的死代码。

但是，如果你真的想跳过这个错误，可以使用[空白标识符（ _ ）：](http://golang.org/doc/effective_go.html#blank)


```go
package main

import (
    "fmt" // imported and not used: "fmt"
)

func main() {
    i := 1 // i declared and not used
}
```

变为

```go
package main

import (
    _ "fmt" // no more error
)

func main() {
    i := 1 // no more error
    _ = i
}
```


如下评论中 kostix 所说，您可以在 FAQ 中找到 Go 团队的官方立场：

> 未使用的变量可能表明存在错误，而未使用的导入只会减慢编译速度。在您的代码树中积累足够的未使用导入，事情可能会变得非常缓慢。因此，Go 不允许这两种情况。

## 回答2


我在两年前学习 Go 语言时遇到了这个问题，所以我声明了自己的函数。

```go
// UNUSED allows unused variables to be included in Go programs
func UNUSED(x ...interface{}) {}
```

然后你可以这样使用它：

```go
UNUSED(x)
UNUSED(x, y)
UNUSED(x, y, z)
```

最大的好处是，你可以把任何东西传递给 UNUSED。

是不是比下面的更好？

```go
_, _, _ = x, y, z
```

那取决于你。

## stackoverflow链接

https://stackoverflow.com/questions/21743841/how-to-avoid-annoying-error-declared-and-not-used
