package main

import "fmt"

func main() {
	fmt.Println(sayHelloWithFilter("arya"))
	fmt.Println(sayHelloWithFilter("sukma"))

	fmt.Println("")

	fmt.Println(sayHelloWithFilter2("arya", spanFilter))
	fmt.Println(sayHelloWithFilter2("sukma", spanFilter))
}

// #1 before
func sayHelloWithFilter(name string) string {
	if name == "sukma" {
		name = "******"
	}

	return "Hello " + name
}

// #2 after
type SpamFilter func(string) string

func sayHelloWithFilter2(name string, filter SpamFilter) string {
	nameFiltered := filter(name)
	return "Hello " + nameFiltered
}

func spanFilter(nama string) (name string) {
	if nama == "sukma" {
		return "*****"
	} else {
		return nama
	}
}
