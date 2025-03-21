# 如何将字符串分配给字节数组

## 问题

我想将字符串分配给字节数组:

```go
var arr [20]byte
str := "abc"
for k, v := range []byte(str) {
  arr[k] = byte(v)
}
```

有别的方法吗？

## 回答1

安全并简单:

```go
[]byte("Here is a string....")
```

## stackoverflow链接

https://stackoverflow.com/questions/8032170/how-to-assign-string-to-bytes-array
