package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.ISOWeek())
	fmt.Println(now.Weekday())
	fmt.Println(now.Clock())
	fmt.Println(now.UTC())

	layout := "2006-01-02"
	timeparse, _ := time.Parse(layout, "2020-11-04")

	fmt.Println(timeparse)
}
