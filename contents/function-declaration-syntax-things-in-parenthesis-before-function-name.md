# 函数声明语法：函数名前的括号内内容

## 问题

对不起，我在问题标题中无法更加具体，但我正在阅读一些 Go 代码，并遇到了这种形式的函数声明：

```go
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
...
}
```

来自 https://github.com/mattermost/platform/blob/master/api/context.go

```go
func (s *GracefulServer) BlockingClose() bool {
...
}
```

来自 https://github.com/braintree/manners/blob/master/server.go

括号中的 (h handler) 和 (s *GracefulServer) 代表什么意思？

## 回答1

这被称为“接收者”。在第一种情况下 (h handler) 它是一个值类型，在第二种情况下 (s *GracefulServer) 它是一个指针。Go 中这种方式与其他语言可能略有不同。然而，接收者类型在大多数面向对象编程中大致像类一样工作。这是你调用方法的对象，就像如果我将一些方法 A 放在某个类 Person 中，那么我需要类型 Person 的实例来调用 A （假设它是一个实例方法而不是静态的！）

这里有一个需要注意的地方是，接收者像其他参数一样被推入调用栈，所以如果接收者是一个值类型，就像在 handler 的情况下，你将在这个你调用方法的对象的副本上工作，这意味着像 h.Name = "Evan" 这样的东西在你返回到调用范围后不会持续存在。因此，任何期望改变接收者状态的东西都需要使用指针或返回修改后的值

这是规范中的相关部分；https://golang.org/ref/spec#Method_sets

## 回答2
它们也被称为接收者。定义它们有两种方式。如果你想修改接收者，可以使用指针，例如：

```go
func (s *MyStruct) pointerMethod() { } // method on pointer
```

如果你不需要修改接收者，你可以将接收者定义为值，例如：

```go
func (s MyStruct)  valueMethod()   { } // method on value
```

这个 [Go playground](http://play.golang.org/p/O0O7Nk1SGF) 示例演示了该概念。

```go
package main

import "fmt"

type Mutatable struct {
    a int
    b int
}

func (m Mutatable) StayTheSame() {
    m.a = 5
    m.b = 7
}

func (m *Mutatable) Mutate() {
    m.a = 5
    m.b = 7
}

func main() {
    m := &Mutatable{0, 0}
    fmt.Println(m)
    m.StayTheSame()
    fmt.Println(m)
    m.Mutate()
    fmt.Println(m)
```

上述程序输出为：

```go
&{0 0}
&{0 0}
&{5 7}
```


## stackoverflow链接

https://stackoverflow.com/questions/34031801/function-declaration-syntax-things-in-parenthesis-before-function-name
