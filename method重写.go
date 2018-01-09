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
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

/* 重写Human的method */
func (e *Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi() //Hi,I am Mark you can call me on 222-222-YYYY
	sam.SayHi()  //Hi,I am Sam,I work at Golang Inc. Call me on 111-888-XXXX
}
