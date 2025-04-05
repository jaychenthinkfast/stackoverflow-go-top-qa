# 如何在 Ubuntu 上设置 GOPATH 环境变量？必须编辑哪个文件？

## 问题


我正在尝试执行 go get ：

```go
go get github.com/go-sql-driver/mysql
```

它会失败，并显示以下错误：

```go
package github.com/go-sql-driver/mysql: cannot download, $GOPATH not set. For more details see: go help gopath
```

当我执行 go env 时，会显示以下 Go 值列表：

```go
ubuntu@ip-xxx-x-xx-x:~$ go env
GOARCH="amd64"
GOBIN=""
GOCHAR="6"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH=""
GORACE=""
GOROOT="/usr/lib/go"
GOTOOLDIR="/usr/lib/go/pkg/tool/linux_amd64"
CC="gcc"
GOGCCFLAGS="-g -O2 -fPIC -m64 -pthread"
CGO_ENABLED="1"
```

明显 GOPATH 没有设置，我该如何设置它以及在哪里设置？

我看到很多提到这个错误的帖子，但没有一个能回答我的问题，需要编辑哪个文件来为这个路径提供值？

## 回答1


### 新方法：Go Modules


自 Go 1.11 以来，您不再需要使用 GOPATH。只需进入您的项目目录，然后执行以下操作一次：

```go
go mod init github.com/youruser/yourrepo
```

* 这样，Go 将在该目录创建一个模块根目录。
* 你可以创建任意数量的模块。

## stackoverflow链接

https://stackoverflow.com/questions/21001387/how-do-i-set-the-gopath-environment-variable-on-ubuntu-what-file-must-i-edit
