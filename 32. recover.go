package main

import "fmt"

func main() {
	runApp(true)

	fmt.Println("arya")
}

// recover adalah fungsi untuk menangkap panic message
// dengan recover, proses panic akan terhenti dan program akan tetap berjalan

// panic dan recover mirip seperti try catch

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error:", message)
	}

	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}
