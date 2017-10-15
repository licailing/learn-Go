package main

import "fmt"

/* 初始化中{}的元素个数不能大于[]中的个数 */
// var balance = [6]int{1, 2, 3, 5}

/* 如果忽略[]中的数字不设置数组大小，Go会根据元素的个数来设置数组的大小 */
// var balance1 = []int{1, 2, 3, 5}

func main() {
	/* 数组：具有相同唯一类型（原始类型：整形、字符串。或自定义类型）的一组且长度固定的数据项序列 */
	var n [10]int

	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

/**
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109
*/
