package main

import (
	"fmt"
	"math"
)

/* 自定义结构体 */
type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

/* 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针 */
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	c1 := Circle{10}
	c2 := Circle{15}
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	fmt.Println("Area of c1 is ", c1.area()) //Area of c1 is  314.1592653589793
	fmt.Println("Area of c2 is ", c2.area()) //Area of c2 is  706.8583470577034
	fmt.Println("Area of r1 is ", r1.area()) //Area of r1 is  24
	fmt.Println("Area of r2 is ", r2.area()) //Area of r2 is  36

}
