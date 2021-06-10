# Go中是否有foreach循环

## 问题
Go语言中是否存在foreach结构？我可以使用for循环切片或数组吗？

## 回答1
https://golang.org/ref/spec#For_range
    
    "for"语句搭配"range"迭代遍历数组、切片、字符串、map或者channel中接收的值。
    迭代中把迭代值赋值给迭代变量，然后执行代码块。

一个例子：
``` 
for index, element := range someSlice {
    // index is the index where we are
    // element is the element from someSlice for where we are
}
```
如果你不关心索引，可以使用 _ :
```
for _, element := range someSlice {
// element is the element from someSlice for where we are
}
```
下划线, _，是[空白标识符](https://golang.org/ref/spec#Blank_identifier) ，匿名占位符。

## stackoverflow链接
https://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go