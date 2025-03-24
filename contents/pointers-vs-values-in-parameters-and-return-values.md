# 返回值还是指针

## 问题

在 Go 语言中，有多种方式返回一个 struct 值或其切片。对于单个值，我见过以下几种：

```go
type MyStruct struct {
    Val int
}

func myfunc() MyStruct {
    return MyStruct{Val: 1}
}

func myfunc() *MyStruct {
    return &MyStruct{}
}

func myfunc(s *MyStruct) {
    s.Val = 1
}
```

我理解这些之间的区别。第一个返回结构体的副本，第二个返回在函数内部创建的结构体值的指针，第三个期望传入一个现有的结构体并覆盖其值。

我看到这些模式在各种上下文中都被使用，我想知道关于这些的最佳实践是什么。何时使用哪种？例如，第一个对于小的结构体可能可以接受（因为开销很小），第二个对于大的结构体。第三个如果你想要极端的内存效率，因为你可以轻松地在调用之间重用单个结构体实例。关于何时使用哪种，有什么最佳实践吗？

类似地，关于切片也有同样的问题：

```go
func myfunc() []MyStruct {
    return []MyStruct{ MyStruct{Val: 1} }
}

func myfunc() []*MyStruct {
    return []MyStruct{ &MyStruct{Val: 1} }
}

func myfunc(s *[]MyStruct) {
    *s = []MyStruct{ MyStruct{Val: 1} }
}

func myfunc(s *[]*MyStruct) {
    *s = []MyStruct{ &MyStruct{Val: 1} }
}
```

再次：这里有哪些最佳实践。我知道切片总是指针，所以返回切片指针没有用。然而，我应该返回结构体值的切片、结构体指针的切片，还是应该将切片指针作为参数传递（[在 Go App Engine API 中使用的模式](https://developers.google.com/appengine/docs/go/datastore/reference#Query.GetAll)）？

## 回答1

* 使用接收者指针的方法很常见；接收者的经验法则是，“[如果有疑问，使用指针](https://go.dev/wiki/CodeReviewComments#receiver-type)。”
* 切片、映射、通道、字符串、函数值和接口值在内部使用指针实现，对这些的指针通常是不必要的。
* 其他情况下，对于大的结构体或需要修改的结构体使用指针，否则[传递值](https://go.dev/wiki/CodeReviewComments#pass-values)，因为通过指针意外地修改事物会让人困惑。

---

应该经常使用指针的一个例子：

* **接收者** 比其他参数更常用指针。方法修改它们被调用的东西，或者命名类型是大型结构体的情况并不少见，因此[建议](https://go.dev/wiki/CodeReviewComments#receiver-type)在罕见情况下才不使用指针。
  * Jeff Hodges 的[复制战斗](https://github.com/jmhodges/copyfighter)工具会自动搜索通过值传递的非微型接收者。

一些你不需要使用指针的情况：

* 代码审查指南建议传递小的结构体如 type Point struct { latitude, longitude float64 } ，甚至可能更大的东西作为值，除非你调用的函数需要能够原地修改它们。
  * 值语义避免了意外改变的情况，即这里的赋值会意外地改变那里的值。
  * 通过值传递小的结构体可以更高效，避免[缓存未命中](https://en.wikipedia.org/wiki/Locality_of_reference)或堆分配。在任何情况下，当指针和值的表现相似时，Go 的做法是选择能够提供更自然语义的选项，而不是榨取每一分速度。
  * 因此，[Go Wiki 的](https://go.dev/wiki/CodeReviewComments#pass-values)代码审查评论页面建议，当结构体较小且可能保持这种状态时，应按值传递。
* 对于**切片**，你不需要传递指针来更改数组的元素。 io.Reader.Read(p []byte) 改变了 p 的字节，例如。这可以说是“将小结构体视为值”的特殊情况，因为内部你传递的是一个称为切片头的小结构（参见 [Russ Cox (rsc)](http://research.swtch.com/godata) 的解释）。同样，你不需要指针来修改映射或进行通道通信。
* 对于**需要重新切片**（更改起始/长度/容量）的情况，内置函数如 append 接受一个切片值并返回一个新的切片。我会模仿这种做法；它避免了别名，返回一个新的切片有助于引起调用者注意可能会分配新的数组，并且对调用者来说很熟悉。
  * 并非总是实际遵循该模式。一些工具，如[数据库接口](https://cloud.google.com/appengine/docs/go/datastore/reference#Get)或[序列化器](http://golang.org/pkg/encoding/json/)，需要向编译时未知类型的切片中追加内容。它们有时会在 interface{} 参数中接受切片的指针。
* 映射、通道、字符串以及函数和接口值，如切片，在内部是引用或包含引用的结构，因此如果您只是想避免复制底层数据，您不需要传递它们的指针。（[rsc 写了一篇关于接口值存储的单独文章](http://research.swtch.com/interfaces)）。
  * 您仍然可能需要在更罕见的情况下传递指针以修改调用者的结构体： [flag.StringVar](http://golang.org/pkg/flag/#StringVar) 因此需要 *string ，例如。

## stackoverflow链接

https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values
