package main

import "fmt"

func main() {
	runApp(0)
}

// defer akan dieksekusi setelah sebuah fungsi selesai dieksekusi
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(val int) {
	// defer akan tetap akan dieksekusi walaupun ada error
	defer logging()

	res := 4 / val

	fmt.Println("Run app")
	fmt.Println("Res", res)
}
