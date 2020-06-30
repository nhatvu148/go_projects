package main

import "fmt"

func main() {
	x := 100
	y := 24

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
}
