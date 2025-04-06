# Go 中处理 JSON POST 请求

## 问题

所以我有以下内容，看起来非常不专业，我自己也在想 Go 有比这更好的库，但我找不到 Go 处理 JSON 数据 POST 请求的示例。它们都是表单 POST。

这里是一个示例请求： curl -X POST -d "{\"test\": \"that\"}" http://localhost:8082/test

这里是代码示例，其中包含了日志：

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type test_struct struct {
    Test string
}

func test(rw http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    log.Println(req.Form)
    //LOG: map[{"test": "that"}:[]]
    var t test_struct
    for key, _ := range req.Form {
        log.Println(key)
        //LOG: {"test": "that"}
        err := json.Unmarshal([]byte(key), &t)
        if err != nil {
            log.Println(err.Error())
        }
    }
    log.Println(t.Test)
    //LOG: that
}

func main() {
    http.HandleFunc("/test", test)
    log.Fatal(http.ListenAndServe(":8082", nil))
}
```

一定有更好的方法，对吧？我只是找不到最佳实践。

## 回答1


请使用 json.Decoder 代替 json.Unmarshal 。

```go
func test(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var t test_struct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    log.Println(t.Test)
}
```


## stackoverflow链接

https://stackoverflow.com/questions/15672556/handling-json-post-request-in-go
