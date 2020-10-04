package helper

import "fmt"

var version string = "1.0"
var Application string

// method init() pada sebuah package akan selalu dieksekusi
func init() {
	Application = "Golang"
	fmt.Println("Init function dipanggil")
}

// access modifier
// jika diawali huruf besar, maka bisa diakses oleh file lain
func SayHello(name string) {
	fmt.Println("Hello", name)
}

// jika diawali huruf kecil, tidak bisa diakses oleh file lain
func sayGoodBye() {
	fmt.Println("Goodbye")
}