# GoLang入门训练

![](images/go-background1.png)


![](images/go-background2.png)


![](images/go-background3.jpg)

## GoLang及GoLand的下载安装
- [GoLang下载](https://golang.google.cn/dl/)
- [GoLand下载](https://www.jetbrains.com/go/)


![](images/go-goland.png)

## GOPATH
- [GOPATH](https://www.jianshu.com/p/cf298a0db3fa)
- [golang中GOPATH的简单理解](https://www.cnblogs.com/lurenq/p/10524647.html)

## Bug记录
1. 初次运行GoLang报错：
    ```text
    Error:run after build is not possible
    main file has non-main package or doesn’t contain main function
    ```
    注意检查包名是不是main函数<br/>
    参考：[参考链接](https://blog.csdn.net/lezeqe/article/details/103321773)
2. 运行同一个包里的两个main()报错：
    ```text
    Error: Package variable contains more than one main function
    Consider using File kind instead
    ```
   把不同的main()拆到不同的包里就行了，反正也是练习模式。


## Go关键词
- break
- default
- func
- interface
- select
- case
- defer
- go
- map
- struct
- chan
- else
- goto
- package
- switch
- const
- fallthrough
- if
- range
- type
- continue
- for
- import
- return
- var

## Go预定义标识符
- append
- bool
- byte
- cap
- close
- complex
- complex64
- complex128
- uint16
- copy
- false
- float32
- float64
- imag
- int
- int8
- int16
- uint32
- int32
- int64
- iota
- len
- make
- new
- nil
- panic
- uint64
- print
- println
- real
- recover
- string
- true
- uint
- uint8
- uintptr

## Go数据类型
- 布尔类型
    - true
    - false
- 数值类型
    - uint8
    - uint16
    - uint32
    - uint64
    - int8
    - int16
    - int32
    - int64
- 字符串类型
- 派生类型
    - 指针类型
    - 数组类型
    - 结构化类型
    - Channel类型
    - 函数类型
    - 切片类型
    - 接口类型
    - Map类型 
    
## Go选择语句select用法

- 每个case都必须是一个通信
- 所有channel表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通信可以进行，它就执行，其他被忽略。
- 如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行
- 否则：
    - 如果有default子句，则执行该语句
    - 如果没有default子句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值

## Go注意事项
1. Go允许自增自减，但不允许这种狗写法：`a = a++`，赞！
2. Go不支持三目运算符，赞！多写点儿阳间的分支语句不香吗？
3. 布尔条件不大打小括号（选择语句）
4. switch语句默认break，不需要自己写；反而不需要哦break的时候需要使用`fallthrough`
5. `:=`用的挺广泛的，要适应
6. `Ctrl+C`可终止命令行死循环
7. 命名为`xxx_test.go`的文件可能被当做测试程序
8. null表示为`nil`
9. 变量有三种类型(依据作用域)：局部变量、全局变量、形式参数
10. Go不支持隐式类型转换，只支持合法的显式类型转换
11. channel编程的操作符`<-`用于指定通道的方向，表示发送或接收(如果未指定方向，则为双向通道)
