# Go中是否支持可选参数

## 问题
Go中是否支持可选参数？或者我是否可以定义同名方法但是不同的参数个数？

## 回答1
Go不支持可选参数[也不支持方法重载](https://golang.org/doc/faq#overloading) ：

    如果不需要类型匹配，方法的调度也就简单了。其他语言的经验告诉我们，拥有同名方法但是不同签名偶有有效，
    但可能在实践中令人困扰与脆弱。在Go的类型系统中只通过名称匹配并保持一致的类型是一种主要的简化决策。

## stackoverflow链接
https://stackoverflow.com/questions/2032149/optional-parameters-in-go