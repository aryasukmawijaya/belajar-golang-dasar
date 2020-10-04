package main

import "fmt"

func main() {
	result := getHello("")

	fmt.Println(result)

	hasil := penjumlahan(20, 40)

	fmt.Println(hasil)
}

func getHello(name string) string {
	if name == "" {
		return "Hello brother!"
	} else {
		return "Hello " + name + "!"
	}
}

func penjumlahan(angka1, angka2 int) int {
	return angka1 + angka2
}
