# 如何转换一个以0终止的字节数组到字符串

## 问题
我需要读取[100]byte来传输一堆字符串数据。

因为不是所有字符串都恰好是100个字符长度，所以字节数组的剩余部分被填充了0。

如果我通过string(byteArray[:])转换[100]byte到string，那么末尾的0将显示为^@^@。

在C语言中，字符串将终止于0，那么Go中转换字节数组到字符串的最佳实践是什么？

## 回答1
将数据读入字节切片时应返回读取的字节数并保存，后续使用它来创建字符串。假设n是读取的字节数，那么可以编码表示如下：
```
s := string(byteArray[:n])
```
如果转换全部数据，可以编码如下：
```
s := string(byteArray[:len(byteArray)])
```
等价于：
```
s := string(byteArray)
```
如果你不知道n,那么你可以用bytes包来发现，假设你的输入没有嵌入其它0字符。
``` 
n := bytes.Index(byteArray, []byte{0})
```
或
```
n := bytes.IndexByte(byteArray, 0)
```

## 回答2
``` 
package main

import "fmt"

func CToGoString(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}

func main() {
    c := [100]byte{'a', 'b', 'c'}
    fmt.Println("C: ", len(c), c[:4])
    g := CToGoString(c[:])
    fmt.Println("Go:", len(g), g)
}
```
输出：
``` 
C:  100 [97 98 99 0]
Go: 3 abc
```

## stackoverflow链接
https://stackoverflow.com/questions/14230145/how-can-i-convert-a-zero-terminated-byte-array-to-string