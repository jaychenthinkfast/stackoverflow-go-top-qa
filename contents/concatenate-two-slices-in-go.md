# Go中如何连接两个切片

## 问题
我尝试合并切片[1, 2]和切片[3, 4]。在Go应该如何实现？

我尝试了：
``` 
append([]int{1,2}, []int{3,4})
```
但是返回：
``` 
cannot use []int literal (type []int) as type int in append
```
然后，[文档](https://golang.org/pkg/builtin/#append) 中表明这么编码是可行的，我错过了什么吗？
``` 
slice = append(slice, anotherSlice...)
```

## 回答1
需要在第二个切片后增加...
``` 
//---------------------------vvv
append([]int{1,2}, []int{3,4}...)
```
和其他可变参数函数一样：
``` 
func foo(is ...int) {
    for i := 0; i < len(is); i++ {
        fmt.Println(is[i])
    }
}

func main() {
    foo([]int{9,8,7,6,5}...)
}
```

## stackoverflow链接
https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go

## 扩展
append()是一个可变参数函数，...可以通过切片方式给append()传入多个参数。