package main

import "fmt"

type Customer struct {
	Name, Hobby string
}

func (customer Customer) sayHi() {
	fmt.Println("Hi! My name is", customer.Name)
}

func main() {
	var arya Customer
	arya.Name = "Arya"

	arya.sayHi()

	fmt.Println(arya)
}
