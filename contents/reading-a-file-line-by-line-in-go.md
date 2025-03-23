# 在 Go 中逐行读取文件

## 问题

我找不到`file.ReadLine` 函数。

如何逐行读取文件？

## 回答1

在 Go 1.1 及更高版本中，最简单的方法是使用 [`bufio.Scanner `](http://golang.org/pkg/bufio/#Scanner)。下面是一个从文件中读取行的简单示例：

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("/path/to/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
```

这是逐行读取 Reader 的最简洁方式。

有一点需要注意：扫描程序将在行长超过 65536 个字符时出错。如果您知道行长大于 64K，请使用 Buffer（） 方法增加扫描仪的容量：

```go
...
scanner := bufio.NewScanner(file)

const maxCapacity int = longLineLen  // your required line length
buf := make([]byte, maxCapacity)
scanner.Buffer(buf, maxCapacity)

for scanner.Scan() {
...
```

## stackoverflow链接

https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
