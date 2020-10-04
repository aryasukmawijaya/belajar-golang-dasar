package main

import (
	"errors"
	"fmt"
)

func pembagianNol(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol!")
	} else {
		hasil := nilai / pembagi
		return hasil, nil
	}
}

func main() {
	pembagian1, err := pembagianNol(10, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pembagian1)
	}

	pembagian2, err := pembagianNol(4, 0)

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(pembagian2)
	}
}
