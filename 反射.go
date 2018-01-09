package main

import "fmt"
import "reflect"

type Person struct {
	name string
	age  int
}

func main() {

	//转化为reflect对象
	var x float64 = 3.4
	v1 := reflect.ValueOf(x)                                      //获取反射值,还可以改变值
	fmt.Println("type:", v1.Type())                               //type: float64
	fmt.Println("kind is float64:", v1.Kind() == reflect.Float64) //kind is float64: true
	fmt.Println("value:", v1.Float())                             //value: 3.4

	var y float64 = 3.4
	v2 := reflect.ValueOf(&y)
	v3 := v2.Elem()
	v3.SetFloat(7.1)
	fmt.Printf("y value is:%v", y) //y value is:7.1
}
