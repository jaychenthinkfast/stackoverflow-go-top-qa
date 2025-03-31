# 从映射中获取键的切片

## 问题
在 Go 中有没有更简单/更优雅的方法来获取映射的键的切片？

目前我正在遍历映射并将键复制到切片中：

```go
i := 0
keys := make([]int, len(mymap))
for k := range mymap {
    keys[i] = k
    i++
}
```
## 回答1
这是一个老问题，但我想说两句。回答2稍微简洁一些，但效率略低。既然你知道它的大小，你甚至不需要使用 append：
```go
keys := make([]int, len(mymap))

i := 0
for k := range mymap {
    keys[i] = k
    i++
}
```
在大多数情况下可能不会有太大差别，但这并不需要更多的工作，在我的测试中（使用包含 1,000,000 个随机 int64 键的映射，然后使用每种方法生成键数组十次），直接分配数组成员比使用 append 快约 20%。

虽然设置容量消除了重新分配，但 append 仍然需要在每次 append 时额外检查是否达到了容量。

## 回答2
```go
package main

func main() {
    mymap := make(map[int]string)
    keys := make([]int, 0, len(mymap))
    for k := range mymap {
        keys = append(keys, k)
    }
}
```
在 Go 中要高效，重要的是要最小化内存分配。

## stackoverflow链接
https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map