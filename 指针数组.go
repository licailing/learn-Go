package main

import "fmt"

const MAX int = 3

func main() {
	/* 局部定义数组 */
	a := []int{10, 100, 200}
	var i int
	/* 指向整形的指针数组 */
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

/**
a[0] = 10
a[1] = 100
a[2] = 200
*/
