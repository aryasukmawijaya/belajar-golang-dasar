package main

import (
	"belajar-golang-dasar/helper"
)

func main() {
	helper.SayHello("arya")

	// error, karena di helper diawali oleh huruf kecil
	// cannot refer to unexported name helper.sayGoodBye
	helper.sayGoodBye()
}