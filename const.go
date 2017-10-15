package main

import "fmt"

func main() {
	//常量定义
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	//并行赋值
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
}

/**
面积为：50
1 false str
*/
