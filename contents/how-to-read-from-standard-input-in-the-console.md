# 我如何在控制台中读取标准输入？

## 问题

我想从命令行读取标准输入，但我的尝试都以程序在提示输入之前退出而告终。我正在寻找 C#中 Console.ReadLine()的等效功能。

这是我目前有的：

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)

    fmt.Println("Enter text: ")
    text2 := ""
    fmt.Scanln(text2)
    fmt.Println(text2)

    ln := ""
    fmt.Sscanln("%v", ln)
    fmt.Println(ln)
}
```

## 回答1


我不确定这段代码出了什么问题

```go
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
fmt.Println(text)
```

它在我的机器上运行正常。然而，对于下一段代码，你需要一个指向你分配输入的变量的指针。尝试将 fmt.Scanln(text2) 替换为 fmt.Scanln(&text2) 。不要使用 Sscanln ，因为它解析的是已经存储在内存中的字符串，而不是从 stdin 解析。如果你想做类似的事情，请将其替换为 fmt.Scanf("%s", &ln)

如果这仍然不起作用，你的问题可能是奇怪的系统设置或一个有缺陷的 IDE。


## stackoverflow链接

https://stackoverflow.com/questions/20895552/how-to-read-from-standard-input-in-the-console
