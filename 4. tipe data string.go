package main

import "fmt"

func main() {
	var firstName string = "Arya"
	var lastName string = "Sukma Wijaya"

	fullName := firstName + " " + lastName

	fmt.Println("first name:", firstName)
	fmt.Println("last name:", lastName)
	fmt.Println("fullname:", fullName)

	fmt.Println("Length:", len(fullName))
	fmt.Println("karakter ke-0:", fullName[0])
}
