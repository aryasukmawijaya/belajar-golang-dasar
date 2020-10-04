package main

import "fmt"

func main() {
	tipe, total := sumAll("Penjumlahan", 10, 10, 10, 10)

	fmt.Println(total)
	fmt.Println(tipe)

	// slice as parameter variadic parameter
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	tipe, total = sumAll("Still penjumlahan", slice...)

	fmt.Println(total)
}

// parameter variadic/varargs hanya bisa ditempatkan diposisi terakhir
func sumAll(tipe string, numbers ...int) (string, int) {
	total := 0
	for _, val := range numbers {
		total += val
	}

	return tipe, total
}
