# GOPATH 和 GOROOT 应该设置什么值？

## 问题



我尝试这样安装 doozer：

```go
 $ goinstall github.com/ha/doozer
```

I get these errors.我遇到了这些错误。

> ```go
> goinstall: os: go/build: package could not be found locally
> goinstall: fmt: go/build: package could not be found locally
> goinstall: io: go/build: package could not be found locally
> goinstall: reflect: go/build: package could not be found locally
> goinstall: math: go/build: package could not be found locally
> goinstall: rand: go/build: package could not be found locally
> goinstall: url: go/build: package could not be found locally
> goinstall: net: go/build: package could not be found locally
> goinstall: sync: go/build: package could not be found locally
> goinstall: runtime: go/build: package could not be found locally
> goinstall: strings: go/build: package could not be found locally
> goinstall: sort: go/build: package could not be found locally
> goinstall: strconv: go/build: package could not be found locally
> goinstall: bytes: go/build: package could not be found locally
> goinstall: log: go/build: package could not be found locally
> goinstall: encoding/binary: go/build: package could not be found locally
> ```

## 回答1


这是我简单的设置：

```go
directory for go related things: ~/programming/go
directory for go compiler/tools: ~/programming/go/go-1.4
directory for go software      : ~/programming/go/packages
```

GOROOT, GOPATH, PATH 设置如下：

```bash
export GOROOT=/home/user/programming/go/go-1.4
export GOPATH=/home/user/programming/go/packages
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

简而言之：

GOROOT 是用于设置  go 安装的 编译器/工具的。
GOPATH 是用于您的 go 项目 / 第三方库（使用 "go get" 下载）。

## stackoverflow链接

https://stackoverflow.com/questions/7970390/what-should-be-the-values-of-gopath-and-goroot
