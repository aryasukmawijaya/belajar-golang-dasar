package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpArya NoKTP = "030001010"
	var statusMarried Married = false

	fmt.Println(noKtpArya)
	fmt.Println(statusMarried)
}
