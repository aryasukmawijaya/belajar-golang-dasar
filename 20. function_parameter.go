package main

import "fmt"

func main() {
	sayHelloTo("Arya", "Sukma")
}

func sayHelloTo(firstName, lastName string) {
	fmt.Println("Hello", firstName, lastName+"!")
}
