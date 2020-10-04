package main

import "fmt"

func main() {
	name := "arya"

	switch name {
	case "arya":
		fmt.Println("hello arya")
	case "sukma":
		fmt.Println("hello!")
	default:
		fmt.Println("hi, kenalan dong")
	}
}
