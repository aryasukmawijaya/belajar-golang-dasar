package main

import "fmt"

func main() {
	fName, lName, _ := getFullName()

	fmt.Println(fName, lName)
}

func getFullName() (string, string, string) {
	return "Arya", "Sukma", "Wijaya"
}
