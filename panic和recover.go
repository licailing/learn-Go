package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func main() {
	fmt.Println("panic before")
	throwsPanic(test)
	fmt.Println("throw panic success")
	test()
	fmt.Println("panic after")
}

/* panic恐慌 */
// 、、恐慌可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组
func test() {
	if user == "" {
		panic("no value for $USER")
	}
	fmt.Println("user=", user)
}

/* recover恢复(只能在defer中使用) */
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}

/*
panic before
throw panic success
panic: no value for $USER

goroutine 1 [running]:
panic(0x4957c0, 0xc042038240)
        E:/go/src/runtime/panic.go:500 +0x1af
main.test()
        E:/go_workspace/src/learn-Go/panic和recover.go:22 +0x15b
main.main()
        E:/go_workspace/src/learn-Go/panic和recover.go:14 +0x150
exit status 2
*/
