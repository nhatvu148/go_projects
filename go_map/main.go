package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)

	emails["Vu"] = "vu@gmail.com"
	emails["Nhat"] = "nhat@gmail.com"

	fmt.Println(emails)

	emails2 := map[string]string{"Nguyen": "nguyen@gmail.com", "Van": "van@gmail.com"}

	fmt.Println(emails2)
}
