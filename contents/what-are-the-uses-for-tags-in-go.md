# Go中标签有什么用

## 问题
在[Go语言规范](https://golang.org/ref/spec#Struct_types) 中提到了标签的简要概述。

    一个字段定义可以伴随一个可选的字符串标签。通过反射接口可见标签，否则忽略。
    
    // A struct corresponding to the TimeStamp protocol buffer.
    // The tag strings define the protocol buffer field numbers.
    struct {
        microsec  uint64 "field 1"
        serverIP6 uint64 "field 2"
        process   string "field 3"
    }
这是一个非常简短的解释，是否可以为我提供这些标签的具体使用情况？

## 回答1
标签可以为字段增加属性信息，可通过反射获取。通常用于结构体字段编码解码时提供转换信息，你可以用它存储任何你想存储的属性信息，不管是用于第三方包还是自定义包。
在 [reflect.StructTag](https://golang.org/pkg/reflect/#StructTag) 文档中提及，按照『约定』标签值是空格分隔的列表对  key:"value"，形如：
``` 
type User struct {
    Name string `json:"name" xml:"name"`
}
```
其中key通常表示后续"value"的包，例如上面json 就代表的 [encoding/json](https://golang.org/pkg/encoding/json/) 包。

如果"value"中有多个信息，一般用逗号分隔。
``` 
Name string `json:"name,omitempty" xml:"name"`
```
通常"value"中'-'表示排除字段处理，在json中就意味着不编码解码该字段。

### 如何使用反射来实现自定义标签？
我们可以使用反射（[reflect包](https://golang.org/pkg/reflect/) ) 来获取结构体字段的标签值。
首先我们需要获取结构体的 [Type](https://golang.org/pkg/reflect/#Type),
然后我们可以通过Type.Field(i int) 或者 Type.FieldByName(name string)查询字段。
这些方法返回的[StructField](https://golang.org/pkg/reflect/#StructField) 描述了结构体字段，
StructField.Tag描述了标签值。

之前提到的『约定』。如果你遵循约定，你可以使用 [StructTag.Get(key string)](https://golang.org/pkg/reflect/#StructTag.Get) 
方法来解析标签值并获取指定key的"value"。这个约定具体在Get()方法中实现。如果你没有遵循这个约定，Get()方法将不能解析你的key:"value"
列表对并且不能支持查找。这也不是什么问题，不过你需要单独实现你的列表对解析逻辑。

Go 1.7 中增加了[StructTag.Lookup()](https://golang.org/pkg/reflect/#StructTag.Lookup) 方法类似于Get()方法，不过可以区分
标签中不含有指定key和标签中含有指定key但是值为空。

让我们来看一个例子：
``` 
type User struct {
    Name  string `mytag:"MyName"`
    Email string `mytag:"MyEmail"`
}

u := User{"Bob", "bob@mycompany.com"}
t := reflect.TypeOf(u)

for _, fieldName := range []string{"Name", "Email"} {
    field, found := t.FieldByName(fieldName)
    if !found {
        continue
    }
    fmt.Printf("\nField: User.%s\n", fieldName)
    fmt.Printf("\tWhole tag value : %q\n", field.Tag)
    fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
}
```
输出：
``` 
Field: User.Name
    Whole tag value : "mytag:\"MyName\""
    Value of 'mytag': "MyName"

Field: User.Email
    Whole tag value : "mytag:\"MyEmail\""
    Value of 'mytag': "MyEmail"
```
GopherCon 2015 有一个关于结构体标签的演示文稿：

[The Many Faces of Struct Tags (slide)](https://github.com/gophercon/2015-talks/blob/master/Sam%20Helman%20%26%20Kyle%20Erf%20-%20The%20Many%20Faces%20of%20Struct%20Tags/StructTags.pdf)
(和[视频](https://www.youtube.com/watch?v=_SCRvMunkdA))

以下是一些常见的标签key列表
* json      - 用于 [encoding/json](https://golang.org/pkg/encoding/json/) 包, 详细内容于 [json.Marshal()](https://golang.org/pkg/encoding/json/#Marshal)
* xml       - 用于 [encoding/xml](https://golang.org/pkg/encoding/xml/) 包, 详细内容于 [xml.Marshal()](https://golang.org/pkg/encoding/xml/#Marshal)
* bson      - 用于 [gobson](https://labix.org/gobson), 详细内容于 [bson.Marshal()](http://godoc.org/gopkg.in/mgo.v2/bson#Marshal)
* protobuf  - 用于 [github.com/golang/protobuf/proto](http://godoc.org/github.com/golang/protobuf/proto), 详细内容于包文档
* yaml      - 用于 [gopkg.in/yaml.v2](https://godoc.org/gopkg.in/yaml.v2) 包, 详细内容于 [yaml.Marshal()](https://godoc.org/gopkg.in/yaml.v2#Marshal)
* db        - 用于 [github.com/jmoiron/sqlx](https://godoc.org/github.com/jmoiron/sqlx) 包; 也用于 [github.com/go-gorp/gorp](https://github.com/go-gorp/gorp) 包
* orm       - 用于 [github.com/astaxie/beego/orm](https://godoc.org/github.com/astaxie/beego/orm) 包, 详细内容于 [Models – Beego ORM](https://beego.me/docs/mvc/model/overview.md)
* gorm      - 用于 [github.com/jinzhu/gorm](https://github.com/jinzhu/gorm) 包, 实例可参考文档: [Models](http://jinzhu.me/gorm/models.html)
* valid     - 用于 [github.com/asaskevich/govalidator](https://github.com/asaskevich/govalidator) 包, 项目页面有实例介绍
* datastore - 用于 [appengine/datastore](https://cloud.google.com/appengine/docs/go/datastore/reference) (谷歌应用引擎平台，数据存储服务), 详细内容于 [Properties](https://cloud.google.com/appengine/docs/go/datastore/reference#hdr-Properties)
* schema    - 用于 [github.com/gorilla/schema](http://godoc.org/github.com/gorilla/schema) 用于HTML表单值填充结构, 详细内容于包文档
* asn       - 用于 [encoding/asn1](https://golang.org/pkg/encoding/asn1/) 包, 详细内容于 [asn1.Marshal()](https://golang.org/pkg/encoding/asn1/#Marshal) 和 [asn1.Unmarshal()](https://golang.org/pkg/encoding/asn1/#Unmarshal)
* csv       - 用于 [github.com/gocarina/gocsv](https://github.com/gocarina/gocsv) 包


## stackoverflow链接
https://stackoverflow.com/questions/10858787/what-are-the-uses-for-tags-in-go