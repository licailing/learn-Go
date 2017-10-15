package main

import "fmt"

func main() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5} //数组

	/* 3种格式
	for init; condition;increment {}
	for condition {}
	for {} 与for(;;)一样
	*/
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为：%d\n", a)
	}

	for a < b { //注意：a为0
		a++
		fmt.Printf("a 的值为：%d\n", a)
	}

	//for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

/**
a 的值为：0
a 的值为：1
a 的值为：2
a 的值为：3
a 的值为：4
a 的值为：5
a 的值为：6
a 的值为：7
a 的值为：8
a 的值为：9
a 的值为：1
a 的值为：2
a 的值为：3
a 的值为：4
a 的值为：5
a 的值为：6
a 的值为：7
a 的值为：8
a 的值为：9
a 的值为：10
a 的值为：11
a 的值为：12
a 的值为：13
a 的值为：14
a 的值为：15
第 0 位 x 的值 = 1
第 1 位 x 的值 = 2
第 2 位 x 的值 = 3
第 3 位 x 的值 = 5
第 4 位 x 的值 = 0
第 5 位 x 的值 = 0
*/
