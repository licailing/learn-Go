package main

import "fmt"

func main() {
	var x interface{} //接口类型

	//判断某个 interface 变量中实际存储的变量类型
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 fun(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或string 型")
	default:
		fmt.Printf("未知型")
	}
	//x 的类型：<nil>
}
