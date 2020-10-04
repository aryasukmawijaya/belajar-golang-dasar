package main

import "fmt"

func main() {
	runApp(true)
}

// panic digunakan untuk menghentikan program, dipanggil ketika terjadi erorr
func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}
