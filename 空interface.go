package main

import "fmt"

func main() {
	/* 空interface 可以保存任意类型数据 */
	var a interface{}
	var i int = 1
	var str string = "licailing"
	a = i
	fmt.Println("a is:", a) //a is: 1
	a = str
	fmt.Println("a is:", a) //a is: licailing
}
