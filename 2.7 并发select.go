package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/* 多个通道，用select监听channel上的数据流动 */
/* select默认是阻塞的，只有当channel上有发送或接收可以进行时才会运行 */
/* 当多个channel都准备好的时候，select是随机选择一个执行的 */
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	/*
		1
		1
		2
		3
		5
		8
		13
		21
		34
		55
		quit
	*/
}
