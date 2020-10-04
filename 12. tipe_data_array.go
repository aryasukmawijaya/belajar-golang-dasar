package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Arya"
	names[1] = "Sukma"
	names[2] = "Wijaya"

	fmt.Println(names)

	var values = [3]int{
		4,
		11,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	// len hanya untuk menghitung panjang array, bukan jumlah data array
}
