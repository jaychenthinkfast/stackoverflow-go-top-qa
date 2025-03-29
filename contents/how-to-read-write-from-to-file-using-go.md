# 如何使用 Go 读取/写入文件?

## 问题
我一直在自学 Go，但在尝试读取和写入普通文件时遇到了难题。

我可以做到 inFile, _ := os.Open(INFILE, 0, 0) ，但实际上获取文件内容没有意义，因为读取函数需要一个 []byte 参数。

```go
func (file *File) Read(b []byte) (n int, err Error)
```

## 回答1
让我们列出所有在 Go 中使用的方法来读取和写入文件，并确保它们与 Go 1 兼容。

因为文件 API 最近有变化，大多数其他答案都不适用于 Go 1。在我看来，它们还缺少 bufio ，这是很重要的。

在以下示例中，我通过从源文件读取并写入目标文件来复制文件。

### 从基础知识开始
```go
package main

import (
    "io"
    "os"
)

func main() {
    // open input file
    fi, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()

    // open output file
    fo, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := fi.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }

        // write a chunk
        if _, err := fo.Write(buf[:n]); err != nil {
            panic(err)
        }
    }
}
```
这里我使用了 os.Open 和 os.Create ，它们是围绕 os.OpenFile 的方便包装器。我们通常不需要直接调用 OpenFile 。

注意处理 EOF。 Read 在每次调用时尝试填充 buf ，如果在填充过程中达到文件末尾，则返回 io.EOF 作为错误。在这种情况下， buf 仍然会保留数据。随后的对 Read 的调用将返回读取的字节数为零，并返回相同的 io.EOF 作为错误。任何其他错误都将导致程序崩溃。

### 使用 bufio
```go
package main

import (
    "bufio"
    "io"
    "os"
)

func main() {
    // open input file
    fi, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // make a read buffer
    r := bufio.NewReader(fi)

    // open output file
    fo, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // make a write buffer
    w := bufio.NewWriter(fo)

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }

        // write a chunk
        if _, err := w.Write(buf[:n]); err != nil {
            panic(err)
        }
    }

    if err = w.Flush(); err != nil {
        panic(err)
    }
}
```
bufio 在这里只是作为一个缓冲区，因为我们没有太多与数据打交道的事情。在大多数其他情况下（尤其是文本文件）， bufio 非常有用，因为它为我们提供了一个方便、灵活的 API 来读取和写入，同时它在幕后处理缓冲。

注意：以下代码是为旧版本的 Go（Go 1.15 及之前版本）准备的。事情已经改变（ ioutil 自 Go 1.16 以来已弃用）。对于新方法，请参阅回答2。
###  使用 ioutil
```go
package main

import (
    "io/ioutil"
)

func main() {
    // read the whole file at once
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    // write the whole body at once
    err = ioutil.WriteFile("output.txt", b, 0644)
    if err != nil {
        panic(err)
    }
}
```
简单得就像吃派一样！但只有在你确定你不是在处理大文件时才使用它。



## 回答2
从 Go 1.16 开始，使用 os.ReadFile 将文件加载到内存中，并使用 os.WriteFile 从内存写入文件（ioutil.ReadFile 现在调用 os.ReadFile 并已弃用）。

注意使用 os.ReadFile ，因为它会将整个文件读入内存。

```go
package main

import "os"
import "log"

func main() {
    b, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    // `b` contains everything your file has.
    // This writes it to the Standard Out.
    os.Stdout.Write(b)

    // You can also write it to a file as a whole.
    err = os.WriteFile("destination.txt", b, 0644)
    if err != nil {
        log.Fatal(err)
    }
}
```

## stackoverflow链接
https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file-using-go