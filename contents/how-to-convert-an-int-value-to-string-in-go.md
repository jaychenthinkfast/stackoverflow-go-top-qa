# Go中如何转换int到string

## 问题
我想将123转化成"123"

## 回答1
使用标准库strconv的[Itoa](https://golang.org/pkg/strconv/#Itoa) 方法
例如：
``` 
package main

import (
    "strconv"
    "fmt"
)

func main() {
    t := strconv.Itoa(123)
    fmt.Println(t)
}
```

## stackoverflow链接
https://stackoverflow.com/questions/10105935/how-to-convert-an-int-value-to-string-in-go

## 扩展
Itoa 等价于 FormatInt(int64(i), 10)。

FormatInt：https://golang.org/pkg/strconv/#FormatInt

func FormatInt(i int64, base int) string

base代表进制，通常十进制数base就是10

例子：
``` 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := int64(-42)

	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)

}
```