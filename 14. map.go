package main

import "fmt"

func main() {
	biodata := map[string]string{
		"nama": "arya",
		"umur": "22",
	}

	fmt.Println(biodata["nama"])

	biodata2 := make(map[string]string)
	biodata2["nama"] = "arya"

	fmt.Println(biodata2)
}
