# 删除使用 go get 安装的包

## 问题

我运行了 go get package 下载了一个包，在了解到我需要设置 GOPATH 否则该包会污染我的根 Go 安装。我该如何删除之前安装的包？

## 回答1


Golang 包可以按照以下方式删除。

```go
go get package@none
```

删除包时将版本部分设置为 none 。

## stackoverflow链接

https://stackoverflow.com/questions/13792254/removing-packages-installed-with-go-get
