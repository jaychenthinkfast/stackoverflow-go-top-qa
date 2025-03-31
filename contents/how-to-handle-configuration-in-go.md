# Go 中如何处理配置

## 问题

对于一个 Go 程序，处理配置参数的首选方式是什么？（在其他情况下可能会使用属性文件或 ini 文件）

## 回答1

JSON 格式对我来说工作得相当好。标准库提供了方法来缩进写入数据结构，因此它非常易于阅读。

JSON 的好处是它相对简单易解析，且易于人类阅读/编辑，同时提供了列表和映射的语义（这可能会非常有用），而许多 ini 类型的配置解析器则不是这样。


示例用法：

**conf.json**:

```go
{
    "Users": ["UserA","UserB"],
    "Groups": ["GroupA"]
}
```


**读取配置的程序**

```go
import (
    "encoding/json"
    "os"
    "fmt"
)

type Configuration struct {
    Users    []string
    Groups   []string
}

file, _ := os.Open("conf.json")
defer file.Close()
decoder := json.NewDecoder(file)
configuration := Configuration{}
err := decoder.Decode(&configuration)
if err != nil {
  fmt.Println("error:", err)
}
fmt.Println(configuration.Users) // output: [UserA, UserB]
```


## 回答2


另一个选项是使用 TOML，这是一种由 Tom Preston-Werner 创建的类似 INI 的格式。我为它构建了一个经过充分测试的 [Go 解析器](https://github.com/BurntSushi/toml)。你可以像这里提出的其他选项一样使用它。例如，如果你有在 something.toml 中的这种 TOML 数据

```go
Age = 198
Cats = [ "Cauchy", "Plato" ]
Pi = 3.14
Perfection = [ 6, 28, 496, 8128 ]
DOB = 1987-07-05T05:45:00Z
```
然后您可以用类似以下方式将其加载到您的 Go 程序中：

```go
type Config struct {
    Age int
    Cats []string
    Pi float64
    Perfection []int
    DOB time.Time
}

var conf Config
if _, err := toml.DecodeFile("something.toml", &conf); err != nil {
    // handle error
}
```

## stackoverflow链接

https://stackoverflow.com/questions/16465705/how-to-handle-configuration-in-go
