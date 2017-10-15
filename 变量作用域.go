package main

import "fmt"

/* 全局变量 */
var g int = 20

func main() {
	/* 局部变量 */
	/* 局部变量可以和全局变量名称可以相同，但是函数体内的局部变量会被优先考虑 */
	var g int = 10

	fmt.Printf("结果： %d\n", g) //结果： 10
}
