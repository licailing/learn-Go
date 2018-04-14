package main

import (
	"fmt"
	//"time"
)

func main() {
	ch := make(chan int)

	go func() {
		<-ch
		fmt.Println(1)
	}()
	ch <- 1 //阻塞，直到被取出
	fmt.Println(2)
}
