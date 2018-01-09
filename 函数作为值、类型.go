package main

import "fmt"

type testInt func(int) bool //声明了一个函数类型

/* 奇数 */
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

/* 偶数 */
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

/* 返回奇数或偶数 */
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice=", slice)
	//slice= [1 2 3 4 5 6 7]
	odd := filter(slice, isOdd) //函数作为值传递
	fmt.Println("Odd elements of slice are:", odd)
	//Odd elements of slice are: [1 3 5 7]
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are:", even)
	//Even elements of slice are: [2 4 6]
}
