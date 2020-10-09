package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])a")

	fmt.Println(regex.MatchString("ara"))
	fmt.Println(regex.MatchString("aRYa"))

	search := regex.FindAllString("ara aba aca arya", 2)
	fmt.Println(search)
}
