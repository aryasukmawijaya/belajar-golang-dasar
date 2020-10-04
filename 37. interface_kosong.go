package main

import "fmt"

func main() {
	var nama interface{} = Ups("arya")
	var umur interface{} = Ups(20)

	// akan terjadi error
	var umur2 int = Ups(20)

	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(Ups(true))
}

// interface{} artinya semua tipe data bisa menjadi result
// begitu juga dengan parameter interface{}
func Ups(foo interface{}) interface{} {
	return foo
}
