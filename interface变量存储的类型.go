package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element
type Person struct {
	name string
	age  int
}

/* 实现fmt String 接口 */
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"Dennis", 70}

	/* 断言 value:= element.(int) 在switch中可以用element.(type) */
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
			//list[0] is an int and its value is 1
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
			//list[1] is an string and its value is hello
		case Person:
			fmt.Printf("list[%d] is an Person and its value is %s\n", index, value)
			//list[2] is an Person and its value is (name: Dennis - age: 70 years)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}
