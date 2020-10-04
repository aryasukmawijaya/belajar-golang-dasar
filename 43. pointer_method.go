package main

import "fmt"

type Man struct {
	Name string
}

// usahakan ketika membuat stuct function menggunakan pointer
// karena tidak menduplikat data/memori
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	arya := Man{"Arya"}
	arya.Married()

	fmt.Println(arya)
}
