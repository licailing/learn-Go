package main

import "fmt"

func main() {
	var a int
	var ptr *int

	/* 指向指针的指针 */
	/* 一个指针变量存放的又是另一个指针变量的地址 */
	var pptr **int

	a = 3000

	ptr = &a    //存放变量a的地址：指向变量a
	pptr = &ptr //存放指针ptr的地址：指向指针ptr

	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

/**
变量 a = 3000
指针变量 *ptr = 3000
指向指针的指针变量 **pptr = 3000
*/
