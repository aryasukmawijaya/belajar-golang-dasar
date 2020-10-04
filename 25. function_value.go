package main

import "fmt"

func main() {
	goodbye := getGoodBye

	fmt.Println(getGoodBye("Arya"))
	fmt.Println(goodbye("Sukma"))
}

func getGoodBye(name string) string {
	return "Goodbye " + name + "!"
}
