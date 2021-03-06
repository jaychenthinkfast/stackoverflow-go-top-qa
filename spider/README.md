# spider

## 功能
* 爬取go标签下投票前100问题输出到top100.md。
* 并根据template.md模板内容输出100个问题的md文件在contents文件夹下， md文件中stackoverflow链接为真实链接。

>template.md 模板内容
```
# 这是一个标题

## 问题
这是一个问题描述

## 回答1
这是回答1的描述

## 回答2
这是回答2的描述

## stackoverflow链接
```

>contents/delete-mapkey-in-go.md 根据模板内容增加真实链接
```
# 这是一个标题

## 问题
这是一个问题描述

## 回答1
这是回答1的描述

## 回答2
这是回答2的描述

## stackoverflow链接
https://stackoverflow.com/questions/1736014/delete-mapkey-in-go
```
  
## 内容格式化
* 为保持阅读体验，请按各问题的md文件格式翻译问题
* md文件中的回答数请在翻译时根据实际投票数等综合判断舍取。 
* 如果需要扩展，请在md文件文末增加
```
## 扩展
这是扩展描述
```

>完整形如
```
# 这是一个标题

## 问题
这是一个问题描述

## 回答1
这是回答1的描述

## 回答2
这是回答2的描述

## stackoverflow链接
https://stackoverflow.com/questions/1736014/delete-mapkey-in-go

## 扩展
这是扩展描述
```


