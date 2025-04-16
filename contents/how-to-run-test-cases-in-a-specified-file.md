# 如何在指定文件中运行测试用例？

## 问题

我的测试用例包分散在多个文件中，如果运行 go test <package_name> ，它会运行包中的所有测试用例。

但是没有必要运行所有这些。有没有办法指定一个文件给 go test 运行，这样它就只会运行该文件中定义的测试用例？

## 回答1

有两种方法。简单的一种是使用 -run 标志并提供匹配你想要运行的测试用例名称的模式。

示例：

```go
$ go test packageName -run NameOfTest
```

请参阅[文档](https://pkg.go.dev/cmd/go#hdr-Testing_flags)获取更多信息。

注意，如果其他测试包含字符串 NameOfTest ，则 -run 标志也可能运行这些测试，因为 -run 标志匹配正则表达式。

因此，为确保仅运行名为 'NameOfTest' 的测试，必须使用正则表达式 ^NameOfTest$ ：

```go
$ go test -run "^NameOfTest$" 
```

另一种方法是命名包含您要运行的测试的特定文件：

```go
$ go test foo_test.go
```

但有一个前提条件。如果：

* foo.go 在 package foo 中。
* foo_test.go 在 package foo_test 中，并导入 'foo'。

如果 foo_test.go 和 foo.go 是同一个包（这是一个常见情况），那么你必须命名所有构建 foo_test 所需的其他文件。在这个例子中就是：

```go
$ go test foo_test.go foo.go
```

我建议使用 -run 模式。或者，在可能的情况下，始终运行所有包测试。

## stackoverflow链接

https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file
