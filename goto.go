package main

import "fmt"

func main() {
	var a int = 10

	/**
	goto语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能
	goto label;
	..
	label:statement;
	*/
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为：%d\n", a)
		a++
	}
}

/**
a的值为：10
a的值为：11
a的值为：12
a的值为：13
a的值为：14
a的值为：16
a的值为：17
a的值为：18
a的值为：19
*/
