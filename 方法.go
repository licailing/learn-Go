package main

import "fmt"

/* 自定义结构体 */
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
	//Area of Circle(c1) =  314
}

/* 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针 */
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
