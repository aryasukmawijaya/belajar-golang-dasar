package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke:", counter)
	}

	names := []string{"Arya", "Sukma", "Wijaya"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, name := range names {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "arya"
	person["title"] = "programmer"

	for i, val := range person {
		fmt.Println(i, "=", val)
	}
}
