# 切片是否包含的方法

## 问题

在 Go 中，有没有不需要遍历切片每个元素就能够判断某个元素是否存在的 类似slice.contains(object) 的方法？

## 回答1

截至 Go 1.21，您可以使用之前提到的实验包提升为 stdlib slices 包。

```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	things := []string{"foo", "bar", "baz"}
	fmt.Println(slices.Contains(things, "foo")) // true
}
```

## stackoverflow链接

https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
