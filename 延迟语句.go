package main

import "fmt"

func main() {
	//延迟语句defer,函数退出前会执行(逆序--后出现的先执行)
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
	/*
		(逆序--后出现的先执行)
			4
			3
			2
			1
			0
	*/
}
