package main

import (
	"fmt"
	"time"
)

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	/* 显式的关闭信道：被关闭的信道会禁止数据流入, 是只读的。我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了。 */
	close(c)
}

func main() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	go counting(bus)
	/* 读取信道：range不等到信道关闭是不会结束读取的 */
	for count := range bus {
		fmt.Println("count:", count)
	}
}

/*
1) GDB调试简介
编译Go程序的时候需要注意几点：
1.传递参数 -ldflags "-s",忽略debug的打印信息
2.传递 -gcflags "-N -l"参数，可以忽略Go内部做的一些优化，聚合变量和函数等优化，
这对于GDB调试来说非常困难，所以在编译的时候加入这两个参数避免这些优化
2) 常用命令
list      	(简写l) --显示源码，默认显示10行代码，后面带上参数显示具体行
break    	(简写b) --设置断点,b 10在第10行设置断点
delete   	(简写d) --删除断点(序号), d 2 删除第二个断点
backtrace 	(简写bt) --打印执行代码的过程
info		--显示信息
    locals			--显示当前执行的程序中的变量值
    breakpoints  	--显示当前设置的断点列表
    goroutines		--显示当前执行的goroutine列表，带*的表示当前执行的
print 		(简写p) --打印变量或其他信息
whatis		--显示当前变量的类型
next		(简写n) --单步调试，跳到下一步，当有断点之后，可以输入n跳转到下一步继续执行
continue	(简写c) --跳出当前断点，后面可以跟参数n，跳过多少次断点
set variable 		--改变运行过程中的变量值,set variable <var>=<value>
*/

/*
1.编译文件，生成可执行文件gdbfile:
go build -gcflags "-N -l" gdbfile.go
2.通过gdb命令启动调试:
gdb gdbfile
*/
