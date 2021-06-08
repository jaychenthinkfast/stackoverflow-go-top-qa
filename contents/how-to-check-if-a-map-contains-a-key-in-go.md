# Go中如何判断一个map包含指定key

## 问题
我知道可以通过遍历map的方式来查找key：
```
for k, v := range m { ... }
```
但是是否有更加高效的方式来检测map中是否存在该key呢？

我在[语言规范](https://golang.org/ref/spec) 里没有找到答案。

## 回答1
单行答案：
``` 
if val, ok := dict["foo"]; ok {
    //do something here
}
```
解释：

Go中的if语句可以同时包含条件语句和初始化语句。上面例子中就同时使用了：
* 初始化两个变量，其中val将从map中接收"foo"的值或者"零值"；而ok将接收一个bool值，如果key "foo"在map中存在则ok会被置为true。
* 判断ok值为true则确定"foo"在map中。

如果map中确实存在"foo"，则将执行if语句的主体，并且val将是该范围的局部变量。

## stackoverflow链接
https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go

## 扩展
在题主的[语言规范](https://golang.org/ref/spec) 中的[索引表达式](https://golang.org/ref/spec#Index_expressions) 部分其实有关于这个问题的解释。
