package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type UserSlice []Person

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[i], value[j]
}

func main() {
	arya := []Person{
		{"arya", 22},
		{"budi", 50},
		{"joko", 30},
	}

	fmt.Println(arya)
	sort.Sort(UserSlice(arya))
	fmt.Println(arya)
}
