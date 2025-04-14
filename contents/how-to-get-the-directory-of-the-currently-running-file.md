# 如何获取当前运行文件的目录？

## 问题


在 nodejs 中我使用__dirname。在 Golang 中它的等价是什么？

我已经谷歌搜索并找到了这篇文章 http://andrewbrookins.com/tech/golang-get-directory-of-the-current-file/。他在下面使用了以下代码

```go
_, filename, _, _ := runtime.Caller(1)
f, err := os.Open(path.Join(path.Dir(filename), "data.csv"))
```

但这是否是 Golang 中正确或惯用的做法？

## 回答1


编辑：自 Go 1.8 版本（2017 年 2 月发布）以来，推荐的做法是使用 os.Executable :

> 可执行文件返回启动当前进程的可执行文件的路径名。无法保证该路径仍然指向正确的可执行文件。如果使用了符号链接来启动进程，则根据操作系统，结果可能是符号链接或它指向的路径。如果需要稳定的结果，可以使用 path/filepath.EvalSymlinks。

要获取可执行文件的目录，可以使用 path/filepath.Dir .

[示例](https://play.golang.org/p/_aolLr7uEH):  

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println(exPath)
}
```


## stackoverflow链接

https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
