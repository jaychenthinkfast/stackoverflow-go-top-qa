# Go中如何检测文件是否存在

## 问题
Go标准库没有仅用于检测文件是否存在的方法（类似于Python的[os.path.exists](https://docs.python.org/3/library/os.path.html#os.path.exists) )。
那么Go的常用方式是什么呢？

## 回答1
检测文件不存在，等价于Python的 if not os.path.exists(filename):
``` 
if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
  // path/to/whatever does not exist
}
```
检测文件存在，等价于Python的if os.path.exists(filename):
```
if _, err := os.Stat("/path/to/whatever"); err == nil {
  // path/to/whatever 存在
} else if os.IsNotExist(err) {
  // path/to/whatever 不存在
} else {
  // 薛定谔: 文件可能存在也可能不存在，需要看err详情
  // 因此，不能使用!os.IsNotExist(err) 判断文件存在
} 
```

## stackoverflow链接
https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go