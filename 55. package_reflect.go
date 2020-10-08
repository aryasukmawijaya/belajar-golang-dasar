package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"true" filetype:"jpg" require:"true"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("require") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}

	return true
}

func main() {
	sample := Person{""}

	fmt.Println(isValid(sample))

	// sampleType := reflect.TypeOf(sample)

	// fmt.Println(sampleType.NumField())

	// fmt.Println(sampleType.Field(0).Name)
	// fmt.Println(sampleType.Field(0).Type)

	// fmt.Println(sampleType.Field(0).Tag.Get("json"))
	// fmt.Println(sampleType.Field(0).Tag.Get("filetype"))
}
