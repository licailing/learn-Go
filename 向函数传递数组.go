package main

import "fmt"

func main() {
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	avg = getAverage(balance, 5)

	fmt.Printf("平均值为：%f", avg)
}

/* 传数组 */
func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum / size)
	return avg //平均值为：214.000000
}
