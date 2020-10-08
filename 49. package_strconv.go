package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err)
	}

	bool_string := strconv.FormatBool(boolean)

	if err == nil {
		fmt.Println(bool_string)
	} else {
		fmt.Println(err)
	}
}
