package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-2222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-2222-XXX"}, "Harvard", 0.00}
	sam := Employee{Human{"Sam", 36, "444-2222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Tings Ltd.", 5000}
	var i Men
	i = mike
	fmt.Println("This is Mike, a Student:") //This is Mike, a Student:
	i.SayHi()                               //Hi,I am Mike you can call me on 222-2222-XXX
	i.Sing("November rain")                 //La la la la... November rain

	i = tom
	fmt.Println("This is Tom, an Employee:") //This is Tom, an Employee:
	i.SayHi()                                //Hi,I am Tom,I work at Tings Ltd.. Call me on 222-444-XXX
	i.Sing("Born to be wild")                //La la la la... Born to be wild

	fmt.Println("Let's use a slice of Men and see what happens") //Let's use a slice of Men and see what happens
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}

	// Hi,I am Paul you can call me on 111-2222-XXX
	// Hi,I am Sam,I work at Golang Inc.. Call me on 444-2222-XXX
	// Hi,I am Mike you can call me on 222-2222-XXX
}
