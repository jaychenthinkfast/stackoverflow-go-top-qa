# 初始化空切片的正确方法

## 问题

声明一个非固定大小的空切片，这样做更好吗：

```go
mySlice1 := make([]int, 0)
```

or:

```go
mySlice2 := []int{}
```

想知道哪一个是正确的方法。

## 回答1

你提供的两种选择在语义上是相同的，但使用 make([]int, 0) 将导致对 runtime.makeslice 的内部调用（Go 1.16）。

你还可以选择保留其值为 nil ：

```go
var myslice []int
```

如 [Golang.org ](http://blog.golang.org/slices)博客中所述：

> 空切片在功能上等同于长度为零的切片，尽管它指向空。它的长度为零，并且可以添加元素，需要分配内存。

然而， nil 切片将 json.Marshal() 进入 "null" ，而空切片将序列化为 "[]" ，正如 @farwayer 所指出的。

如 @ArmanOrdookhani 所指出的，上述任何选项都不会引起任何内存分配。

## 回答2


它们是等价的。请看以下代码：

```go
mySlice1 := make([]int, 0)
mySlice2 := []int{}
fmt.Println("mySlice1", cap(mySlice1))
fmt.Println("mySlice2", cap(mySlice2))
```

输出:

```go
mySlice1 0
mySlice2 0
```

两个切片都有 0 容量，这意味着两个切片的长度（不能大于容量）相同，这也就意味着两个切片没有元素。这意味着两个切片在各个方面都是相同的。

## stackoverflow链接

https://stackoverflow.com/questions/29164375/correct-way-to-initialize-empty-slice
