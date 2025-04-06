# 在 Go 中将字符串转换为整型？

## 问题
我正在尝试将来自 flag.Arg(n) 的字符串转换为 int 。在 Go 中这样做有什么惯用方法？

## 回答1
例如 strconv.Atoi 。

代码：

```go
package main

import (
"fmt"
"strconv"
)

func main() {
s := "123"

    // string to int
    i, err := strconv.Atoi(s)
    if err != nil {
        // ... handle error
        panic(err)
    }

    fmt.Println(s, i)
}
```

## stackoverflow链接
https://stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go