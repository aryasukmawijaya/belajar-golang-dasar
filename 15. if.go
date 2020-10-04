package main

import "fmt"

func main() {
	var nama string = "arya sukma"

	if nama == "arya" {
		fmt.Println("hello arya")
	} else {
		fmt.Println("hi! kenalan dong")
	}

	if panjangNama := len(nama); panjangNama > 4 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
