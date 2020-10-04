package main

import "fmt"

type Address struct {
	City, Country string
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	jakarta := &Address{"Jakarta", ""}
	changeAddressToIndonesia(jakarta)

	fmt.Println(jakarta)

	bandung := Address{"Bandung", ""}
	changeAddressToIndonesia(&bandung)

	fmt.Println(bandung)
}
