package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.Int("host", 1, "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "", "Put your database password")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
}
