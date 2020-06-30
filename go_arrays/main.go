package main

import "fmt"

func main() {
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	fmt.Println(fruitArr)

	// fruit2 := [2]string{"Orange", "Mango"}
	fruit3 := []string{"Orange", "Mango", "Banana"}

	fmt.Println(len(fruit3))
	fmt.Println(fruit3[1:3])
}
