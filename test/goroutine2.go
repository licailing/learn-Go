package main

import (
	"fmt"
	"time"
)

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

func consumer(p <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-p
		fmt.Println("receive:", v)
	}
}

func main() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	time.Sleep(time.Second*1)
}
