# Go中如何找到对象的类型

## 问题
Go中如何找到对象的类型?在Python中可以用typeof来获取对象的类型。
同样在Go中，是否有类似的方法？

## 回答1
Go的reflect包有方法用于检测变量类型。
``` 
package main

import (
	"fmt"
	"reflect"
)

func main() {
	tst := "string"
	tst2 := 10
	tst3 := 1.2

	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))

}
```
输出：
``` 
string
int
float64
```
可以在https://play.golang.org/p/XQMcUVsOja 中查看执行情况。

更多说明可参考：http://golang.org/pkg/reflect/#Type

## 回答2
这是回答2的描述

## stackoverflow链接
https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go