package main

import "fmt"

func main() {
	// var name = "Nhat Vu"
	var age int32 = 30
	var isCool = true

	name := "Vu" // Short hand

	fmt.Println(name, age)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
}
