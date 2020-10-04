package main

// jika melakukan import namun tidak digunakan, agar tidak terjadi error maka bisa menambahkan blank identifier (_) sebelum nama package
import (
	_ "belajar-golang-dasar/helper"
)

func main() {
	// helper.SayHello("arya")

	// error, karena di helper diawali oleh huruf kecil
	// cannot refer to unexported name helper.sayGoodBye
	// helper.sayGoodBye()
}