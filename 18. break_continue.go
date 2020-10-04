package main

import "fmt"

func main() {
	// break menghentikan seluruh perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke", i)
	}

	// continue hanya menghentikan perulangan yang sedang berjalan,
	// dan melanjutkan ke perulangan selanjutnya
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke", i)
	}
}
