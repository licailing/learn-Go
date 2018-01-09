package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)

	fmt.Printf("最大值是：%d\n", ret) //最大值是：200

	m, n := swap("Mahesh", "Kumar")
	fmt.Println(m, n) //Kumar Mahesh

	var x int = 100
	var y int = 200
	fmt.Printf("转换前：x = %d , y = %d\n", x, y)
	swap1(&x, &y) //传址
	fmt.Printf("转换后：x = %d , y = %d\n", x, y)
	/**
	转换前：x = 100 , y = 200
	转换后：x = 200 , y = 100
	*/

	/* 声明函数变量：函数作为值 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	/* 使用函数 */
	fmt.Println(getSquareRoot(9)) //3
}

func max(num1, num2 int) int {
	//声明局部变量
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* 多返回值 */
func swap(x, y string) (string, string) {
	return y, x
}

/* 引用传递指针参数传递到函数内 */
func swap1(x1 *int, y1 *int) {
	var temp int

	temp = *x1
	*x1 = *y1
	*y1 = temp
}

/* 命名返回值 */
func SumAndProcuct(A, B int) (add int, Multipled int) {
	add = A + B
	Multipled = A * B
	return
}
