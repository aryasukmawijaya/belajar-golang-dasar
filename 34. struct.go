package main

import "fmt"

type Customer struct {
	Name, Address string
	Hobby         string
	Age           int
	Married       bool
}

func main() {
	var customer1 Customer
	customer1.Name = "Arya"
	customer1.Address = "Indonesia"
	customer1.Hobby = "Coding"
	customer1.Age = 22

	fmt.Println(customer1)

	customer2 := Customer{
		Name:    "sukma",
		Address: "indonesia",
		Hobby:   "coding",
		Age:     23,
	}

	// struct literal
	customer3 := Customer{"wijaya", "indonesia", "coding", 24, true}
	fmt.Println(customer2)
	fmt.Println(customer3)
}
