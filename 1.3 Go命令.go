package main

/*
**Go命令
1> go build [文件.go]
编译代码，默认编译当前目录下所有go文件，在包的编译过程中，若有必要，会同时编译与之相关联的包

2> go clean
移除当前源码包和关联源码包里面编译生成的文件

-i 清除关联的安装的包和可运行文件，也就是通过go install安装的文件
-n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的
-r 循环的清除在import中引入的包
-x 打印出来执行的详细命令，其实就是-n打印的执行版本

3> go fmt <文件名>.go
其实调用gofmt，而且需要参数-w，否则格式化结果不会写入文件。gofmt -w -l src，可以格式化整个项目。

gofmt的参数介绍

-l 显示那些需要格式化的文件
-w 把改写后的内容直接写入到文件中，而不是作为结果打印到标准输出。
-r 添加形如“a[b:len(a)] -> a[b:]”的重写规则，方便我们做批量替换
-s 简化文件中的代码
-d 显示格式化前后的diff而不是写入文件，默认是false
-e 打印所有的语法错误到标准输出。如果不使用此标记，则只会打印不同行的前10个错误。
-cpuprofile 支持调试模式，写入相应的cpufile到指定的文件

4> go get
下载和安装依赖包

5> go install
第一步是生成结果文件(可执行文件或者.a包)，第二步会把编译好的结果移到$GOPATH/pkg或者$GOPATH/bin。

6> go test
执行这个命令，会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件。

-bench regexp 执行相应的benchmarks，例如 -bench=.
-cover 开启测试覆盖率
-run regexp 只运行regexp匹配的函数，例如 -run=Array 那么就执行包含有Array开头的函数
-v 显示测试的详细命令

> go version
查看go当前的版本

> go env
查看当前go的环境变量

> go list
列出当前全部安装的package

> go run
编译并运行Go程序

*/

