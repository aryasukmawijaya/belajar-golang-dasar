package main

import "fmt"

func main() {
	nama, umur, titel := getPersonDetail()

	fmt.Println(nama, umur, titel)
	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(titel)
}

func getPersonDetail() (fullName string, age int, title string) {
	namaLengkap := "Arya Sukma Wijaya"
	age = 22
	title = "Programmer"

	return namaLengkap, age, title
}
