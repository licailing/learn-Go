package main

import "fmt"

func main() {
	//定义局部变量
	var grade string = "B"
	var marks int = 90

	//不需要加break
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70: //同时测试多个可能符合条件的值
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是%s\n", grade)

	score := 98
	switch {
	case score > 90:
		fmt.Printf(">90")
	case score > 80:
		fmt.Printf(">80")
	case score > 70:
		fmt.Printf(">70")
	}

}

/**
优秀
你的等级是A
>90
*/
