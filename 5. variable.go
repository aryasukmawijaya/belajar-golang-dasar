package main

import "fmt"

func main() {
	var firstName string
	firstName = "Arya"

	var lastName = "Sukma"
	age := 20

	fmt.Println("first name:", firstName)
	fmt.Println("last name:", lastName)
	fmt.Println("umur:", age)

	age = 19
	fmt.Println("umur:", age)

	var namaDepan, namaBelakang = "arya", "sukma wijaya"

	fmt.Println("nama depan:", namaDepan)
	fmt.Println("nama belakang:", namaBelakang)
}
