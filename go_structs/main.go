package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// func (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " and I am " + strconv.Itoa(p.age)
}

// func (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	person1 := Person{"Nhat", "Vu", "Long Thanh", "Male", 30}

	fmt.Println(person1)

	person1.hasBirthday()
	fmt.Println(person1.greet())
}
