package main

import "fmt"

//返回一个函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()
	fmt.Println(nextNumber()) //1
	fmt.Println(nextNumber()) //2
	fmt.Println(nextNumber()) //3

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1()) //1
	fmt.Println(nextNumber1()) //2
}
