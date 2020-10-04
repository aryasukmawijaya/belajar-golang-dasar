package main

import "fmt"

func main() {
	name := "arya"
	counter := 0

	// sebuah fungsi bisa mengakses data dalam scopenya
	increment := func() {
		// lakukan deklarasi ulang didalam variabel agar tidak mengubah variabel diluar scope
		name := "sukma"

		fmt.Println("Name:", name)
		counter++

		decrement := func() {
			fmt.Println("Decrement")
			counter--
		}

		decrement()
	}

	// fungsi decrement tidak bisa diakses karena diluar scope main
	// decrement()

	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
