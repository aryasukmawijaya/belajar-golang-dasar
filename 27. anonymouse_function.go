package main

import "fmt"

func main() {
	blockedAdmin := func(name string) bool {
		return name == "admin"
	}

	blockedRoot := func(name string) bool {
		return name == "root"
	}

	registerUser("admin", blockedAdmin)
	registerUser("arya", blockedRoot)
	registerUser("sukma", func(name string) bool {
		return name == "sukma"
	})
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked,", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
