package main

import "fmt"

func main() {
	var result = 10 + 10

	fmt.Println(result)

	var a = 4
	var b = 9

	fmt.Println(a * b)

	// augmented assignment
	a += 100

	fmt.Println(a)

	// unary operator
	a++

	fmt.Println(a)

	var married = false
	var status = !married
	fmt.Println(status)
}
