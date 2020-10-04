package main

import "fmt"

func main() {
	// assertions adalah konvert tipe data
	// tipe data res adalah interface{}, walaupun datanya adalah string
	res := random()

	// convert ke string
	// resString := res.(string)

	// coba convert ke int
	// resInteger := res.(int) // panic

	// fmt.Println(resString)

	// mmenggunakan switch
	switch val := res.(type) {
	case string:
		fmt.Println("String:", val)
	case int:
		fmt.Println("Integer:", val)
	case bool:
		fmt.Println("Boolean:", val)
	default:
		fmt.Println("Unknown")
	}
}

func random() interface{} {
	return true
}
