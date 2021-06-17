# 如何在控制台中打印结构变量

## 问题
如何在控制台中打印这个结构变量Id, Title, Name等。
```
type Project struct {
    Id      int64   `json:"project_id"`
    Title   string  `json:"title"`
    Name    string  `json:"name"`
    Data    Data    `json:"data"`
    Commits Commits `json:"commits"`
}
```

## 回答1
打印结构体字段名:
``` 
fmt.Printf("%+v\n", yourProject)
```
在[fmt包](https://golang.org/pkg/fmt/) 里：

    在打印结构体时，%+v 将增加打印字段名

基于假设你有一个Project的实例yourProject

## 回答2
推荐[go-spew](https://github.com/davecgh/go-spew) :实现了一个深度美观的打印机
用于Go数据结构的调试。
``` 
go get -u github.com/davecgh/go-spew/spew
```
实例：
``` 
package main

import (
    "github.com/davecgh/go-spew/spew"
)

type Project struct {
    Id      int64  `json:"project_id"`
    Title   string `json:"title"`
    Name    string `json:"name"`
    Data    string `json:"data"`
    Commits string `json:"commits"`
}

func main() {

    o := Project{Name: "hello", Title: "world"}
    spew.Dump(o)
}
```
输出：
``` 
(main.Project) {
 Id: (int64) 0,
 Title: (string) (len=5) "world",
 Name: (string) (len=5) "hello",
 Data: (string) "",
 Commits: (string) ""
}
```



## stackoverflow链接
https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console