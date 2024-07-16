package main

import "fmt"

func main() {
	x := factorial(5)
	fmt.Println(x)
}

// Recursive function example
func factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}
