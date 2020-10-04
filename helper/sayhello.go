package helper

import "fmt"

var version string = "1.0"
var Application string = "Golang"

// access modifier
// jika diawali huruf besar, maka bisa diakses oleh file lain
func SayHello(name string) {
	fmt.Println("Hello", name)
}

// jika diawali huruf kecil, tidak bisa diakses oleh file lain
func sayGoodBye() {
	fmt.Println("Goodbye")
}