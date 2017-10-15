package main

import "fmt"

func main() {
	var a int = 10
	/* 取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址 */
	fmt.Printf("a的地址为：%x\n", &a)

	/* var var_name *var-type(指针类型-指向类型) */
	var ip *int /* 指向整形 */
	ip = &a
	fmt.Printf("ip的值为：%x\n", ip)
	fmt.Printf("ip的地址为：%x\n", &ip)
	fmt.Printf("ip指向的值为:%d\n", *ip)
}

/**
a的地址为：114de0b0
ip的值为：114de0b0
ip的地址为：114d4108
ip指向的值为:10
*/
