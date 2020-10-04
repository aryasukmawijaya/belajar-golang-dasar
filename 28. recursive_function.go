package main

import "fmt"

func main() {
	// loop := factorialLoop(4)
	// factorialRecursive(4)
	// fmt.Println(loop)
	// fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorialRecursive(4))
}

func factorialLoop(value int) (result int) {
	result = 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		value *= factorialRecursive(value - 1)
	}

	return value
}
