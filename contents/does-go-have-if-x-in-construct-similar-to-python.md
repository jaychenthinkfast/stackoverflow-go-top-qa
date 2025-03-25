# Go 是否有类似于 Python 的 "if x in" 构造？

## 问题


如何在不遍历整个数组的情况下检查 x 是否在数组中，使用 Go？该语言是否有这样的构造？

就像在 Python 中：

```python
if "x" in array: 
  # do something
```


## 回答1


从 Go 1.18 或更高版本开始，您可以使用 slices.Contains 。

在 Go 1.18 之前没有内置操作符。您需要遍历数组。您必须编写函数来完成它，如下所示：

```go
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
```

如果您想能够在不遍历整个列表的情况下检查成员资格，则需要使用映射而不是数组或切片，如下所示：

```go
visitedURL := map[string]bool {
    "http://www.google.com": true,
    "https://paypal.com": true,
}
if visitedURL[thisSite] {
    fmt.Println("Already been here.")
}
```


## stackoverflow链接

https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python
