# 如何将持续时间乘以整数？

## 问题


为了测试并发 goroutine，我在一个函数中添加了一行代码，使其返回随机时间（最多一秒）

```go
time.Sleep(rand.Int31n(1000) * time.Millisecond)
```

然而，当我编译时，我遇到了这个错误

> .\\crawler.go:49: invalid operation: rand.Int31n(1000) \* time.Millisecond (mismatched types int32 and time.Duration).

有什么办法吗？我该如何乘以一个持续时间？

## 回答1


int32 和 time.Duration 是不同类型。您需要将 int32 转换为 time.Duration ：

```go
time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
```

## stackoverflow链接

https://stackoverflow.com/questions/17573190/how-to-multiply-duration-by-integer
