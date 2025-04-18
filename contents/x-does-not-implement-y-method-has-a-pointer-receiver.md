# Golang 方法使用指针接收器

## 问题


我有这个示例代码

```go
package main

import (
    "fmt"
)

type IFace interface {
    SetSomeField(newValue string)
    GetSomeField() string
}

type Implementation struct {
    someField string
}

func (i Implementation) GetSomeField() string {
    return i.someField
}

func (i Implementation) SetSomeField(newValue string) {
    i.someField = newValue
}

func Create() IFace {
    obj := Implementation{someField: "Hello"}
    return obj // <= Offending line
}

func main() {
    a := Create()
    a.SetSomeField("World")
    fmt.Println(a.GetSomeField())
}
```

SetSomeField 不按预期工作，因为它的接收者不是指针类型。

如果我将方法改为指针接收者，我期望能正常工作，看起来是这样的：

```go
func (i *Implementation) SetSomeField(newValue string) { ...
```

编译此代码会导致以下错误：

```go
prog.go:26: cannot use obj (type Implementation) as type IFace in return argument:
Implementation does not implement IFace (GetSomeField method has pointer receiver)
```

我该如何让 struct 实现接口，同时让方法 SetSomeField 改变实际实例的值而不创建副本？


## 回答1


你的结构体指针应该实现该接口。这样你就可以修改它的字段。

看我如何修改你的代码，使其按你期望的方式工作：

```go
package main

import (
    "fmt"
)

type IFace interface {
    SetSomeField(newValue string)
    GetSomeField() string
}

type Implementation struct {
    someField string
}  

func (i *Implementation) GetSomeField() string {
    return i.someField
}

func (i *Implementation) SetSomeField(newValue string) {
    i.someField = newValue
}

func Create() *Implementation {
    return &Implementation{someField: "Hello"}
}

func main() {
    var a IFace
    a = Create()
    a.SetSomeField("World")
    fmt.Println(a.GetSomeField())
}
```


## stackoverflow链接

https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
