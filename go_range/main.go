package main

import "fmt"

func main() {
	ids := []int{12, 45, 32, 67, 2}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	emails := map[string]string{"Nguyen": "nguyen@gmail.com", "Van": "van@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
