package main

import "fmt"

type Person struct {
	Name string
}

func sayHi(person Person) interface{} {
	if person.Name == "" {
		return nil
	}

	return "Hi " + person.Name
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			name: name,
		}
	}
}

func main() {
	arya := Person{"sukma"}
	hi := sayHi(arya)

	if hi != nil {
		fmt.Println(hi)
	}

	nama := newMap("sukma")
	if nama != nil {
		fmt.Println(nama)
	}
}
