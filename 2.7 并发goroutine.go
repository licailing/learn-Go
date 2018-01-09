package main

import (
	"fmt"
	"runtime"
)

/* goroutine 协程 */

func say(s string) {
	for i := 0; i < 5; i++ {
		/* 让CPU把时间片让给别人，下次某个时候继续回复执行该goroutine */
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	/* 开启 goroutine */
	go say("world")
	say("hello")
}

/*
hello
world
hello
world
hello
world
hello
world
hello
*/
