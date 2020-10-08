package main

import (
	"fmt"
	"strings"
)

func main() {
	var fullname string = "arya sukma wijaya"

	fmt.Println(strings.Trim(fullname, "a"))
	fmt.Println(strings.Contains(fullname, "sukma"))
	fmt.Println(strings.Split(fullname, " "))
	fmt.Println(strings.ToUpper(fullname))
	fmt.Println(strings.ToLower("ARYA SUKMA WIJAYA"))
	fmt.Println(strings.ToTitle(fullname))
	fmt.Println(strings.ReplaceAll(fullname, "arya", "aryaa"))
	fmt.Println(strings.Replace("arya arya arya arya arya", "arya", "aryaa", 2))
}
