# 遍历映射中的所有键

## 问题


有没有一种方法可以获取 Go 语言映射中所有键的列表？元素的数量由 len() 给出，但如果我有一个像这样的映射：

```go
m := map[string]string{ "key1":"val1", "key2":"val2" };
```

我该如何遍历所有键？

## 回答1


```go
for k, v := range m { 
    fmt.Printf("key[%s] value[%s]\n", k, v)
}
```

or

```go
for k := range m {
    fmt.Printf("key[%s] value[%s]\n", k, m[k])
}
```

[Go 语言规范](http://golang.org/ref/spec#For_statements) 中对于 for 语句指定，第一个值是键，第二个变量是值，但不必一定存在。

## 回答2


这里有一些简单的方法来获取映射中 keys 的 切片 。

```go
// Return keys of the given map
func Keys(m map[string]interface{}) (keys []string) {
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

// use `Keys` func
func main() {
    m := map[string]interface{}{
        "foo": 1,
        "bar": true,
        "baz": "baz",
    }
    fmt.Println(Keys(m)) // [foo bar baz]
}
```


## stackoverflow链接

https://stackoverflow.com/questions/1841443/iterating-over-all-the-keys-of-a-map
