package test

import "fmt"

var quit chan int //只开一个信道
var buffer_quit chan int
var quit_test = make(chan int, 10)

func foo(id int) {
	fmt.Println("sent ", id, " to quit")
	quit <- id
}

func foo1(id int) {
	fmt.Println("sent ", id, " to buffer_quit")
	buffer_quit <- id
}

func foo2(id int) {
	fmt.Println("sent ", id, " to quit_test")
	quit_test <- id
}

func NoBuffer_goRoutines() {
	count := 10
	quit = make(chan int) //无缓冲

	for i := 0; i < count; i++ {
		go foo(i)
	}

	for i := 0; i < count; i++ {
		fmt.Println("recived ", <-quit, " from quit")
	}
}

func Buffer_goRoutines() {
	count := 10
	buffer_quit = make(chan int, count) //缓冲

	for i := 0; i < count; i++ {
		go foo1(i)
	}

	for i := 0; i < count; i++ {
		fmt.Println("recived ", <-buffer_quit, " from buffer_quit")
	}
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("recived ", <-quit_test, " from quit_test")
	}
}

func GoRoutines() {
	for i := 0; i < 10; i++ {
		go foo2(i)
	}
	go test()

	fmt.Println("test")
}
