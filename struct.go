package main

import "fmt"

/* 字符串切片 */
type Skills []string

type Human struct {
	name   string
	age    int
	weight int
	phone  string //Human类型拥有的字段
}

type Student struct {
	Human      //匿名字段,struct
	Skills     //匿名字段,自定义类型string slice
	int        //内置类型作为匿名字段
	speciality string
	phone      string //重写Human中的phone
}

func main() {
	jane := Student{Human: Human{"Jane", 35, 100, "777-444-xxxx"}, speciality: "Biology", phone: "333-222"}
	fmt.Println("Her name is ", jane.name) //Her name is  Jane

	fmt.Println("Her age is ", jane.age)                    //Her age is  35
	fmt.Println("Her weight is ", jane.weight)              //Her weight is  100
	fmt.Println("Her speciality is ", jane.speciality)      //Her speciality is  Biology
	fmt.Println("Her work  phone is ", jane.phone)          //Her work  phone is  333-222
	fmt.Println("Her personal phone is ", jane.Human.phone) //Her personal phone is  777-444-xxxx

	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills) //Her skills are  [anatomy]
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills are ", jane.Skills) //Her skills are  [anatomy physics golang]

	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int) //Her preferred number is  3
}
