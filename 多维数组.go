package main

import "fmt"

/* 多维数组声明方式 */
/* var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type */

func main() {
	/* 二维数组 */
	var a = [5][2]int{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
		{4, 8},
	}

	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

/**
a[0][0] = 0
a[0][1] = 0
a[1][0] = 1
a[1][1] = 2
a[2][0] = 2
a[2][1] = 4
a[3][0] = 3
a[3][1] = 6
a[4][0] = 4
a[4][1] = 8
*/
