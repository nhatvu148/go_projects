package main

import "fmt"

func main() {
	a := 7
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)
}
