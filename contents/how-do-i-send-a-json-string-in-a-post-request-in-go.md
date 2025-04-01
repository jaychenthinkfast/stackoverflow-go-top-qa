# 如何在 Go 中发送 JSON 字符串的 POST 请求

## 问题

我尝试使用 Apiary 并创建了一个通用的模板来向模拟服务器发送 JSON，并编写了以下代码：

```go
package main   
import (
        "encoding/json"
        "fmt"
        "github.com/jmcvetta/napping"
        "log"
        "net/http"
)
func main() {
  url := "http://restapi3.apiary.io/notes"
  fmt.Println("URL:>", url)
  
  s := napping.Session{}
  h := &http.Header{}
  h.Set("X-Custom-Header", "myvalue")
  s.Header = h  
  var jsonStr = []byte(`{ "title": "Buy cheese and bread for breakfast."}`)  
  var data map[string]json.RawMessage
  err := json.Unmarshal(jsonStr, &data)
  if err != nil {
        fmt.Println(err)
  }
  
  resp, err := s.Post(url, &data, nil, nil)
  if err != nil {
        log.Fatal(err)
  }
  fmt.Println("response Status:", resp.Status())
  fmt.Println("response Headers:", resp.HttpResponse().Header)
  fmt.Println("response Body:", resp.RawText())
}
```

这段代码没有正确发送 JSON，但我不知道为什么。每次调用的 JSON 字符串都可能不同。我无法使用 Struct 来实现这一点。

## 回答1


我不熟悉 napping ，但使用 Golang 的 net/http 包运行良好（ playground）：

```go
func main() {
    url := "http://restapi3.apiary.io/notes"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := io.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
```


## stackoverflow链接

https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
