package main

import "fmt"

func main() {
	var i int = 10
	var m int32
	var n float32
	var ptr *int //指针类型

	fmt.Printf("第1行 - i 变量类型为 = %T\n", i)
	fmt.Printf("第2行 - m 变量类型为 = %T\n", m)
	fmt.Printf("第3行 - n 变量类型为 = %T\n", n)

	ptr = &i //返回变量的存储地址
	fmt.Printf("i 的值为 %d\n", i)
	fmt.Printf("ptr 为 %d\n", ptr)
	fmt.Printf("ptr 为 %d\n", *ptr)
}

/**
第1行 - i 变量类型为 = int
第2行 - m 变量类型为 = int32
第3行 - n 变量类型为 = float32
i 的值为 10
ptr 为 291938432
ptr 为 10
*/
