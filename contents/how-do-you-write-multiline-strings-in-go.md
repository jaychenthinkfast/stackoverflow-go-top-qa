# Go中如何编写多行字符串

## 问题
类似于Python中多行字符串的写法，Go中是否有相似的实现：
``` 
"""line 1
line 2
line 3"""
```
如果没有，是否有其他跨多行字符串的推荐方案？

## 回答1
根据[Go语言规范](https://golang.org/ref/spec#String_literals) 你可以使用原生字符串字面量, 其中字符串使用反引号而不是双引号。
```
`line 1
line 2
line 3`
```

## 回答2 
你可以这样写：
```
"line 1" +
"line 2" +
"line 3"
```
等同于
```
"line 1line 2line 3"
```
不同于反引号的方式，它将解析转义字符。需要注意的是"+"必须跟着先导行，下面的写法就是错误的：
``` 
"line 1"
+"line 2"
```

## stackoverflow链接
https://stackoverflow.com/questions/7933460/how-do-you-write-multiline-strings-in-go