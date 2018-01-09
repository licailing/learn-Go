package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	/* 发送total到channel c */
	/* 任何发送都会被阻塞，直到数据被读出 */
	c <- total
}



func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	/* 关闭channel */
	close(c)
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	/* 无缓存channel */
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	/* 从channel c中接收数据，并赋值给x,y */
	/* 读取将会被阻塞，直到有数据接收 */
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) //17 -5 12

	/* 缓存channel */
	c_b := make(chan int, 10)
	/* cap:数组切片分配的空间大小 */
	go fibonacci(cap(c_b), c_b)
	/* 能够不断的读取channel中的数据，直到channel被显式的关闭 */
	for i := range c_b {
		fmt.Println(i)
	}
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
	*/
}
