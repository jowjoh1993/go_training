package main

import "fmt"

func main() {
	// Call variadic function
	sum := sumup(1, 10, 15, 20, -5)
	fmt.Println(sum)

	numbers := []int{1, 10, 15, 23}

	// Split slice into parameter values
	anotherSum := sumup(numbers...)
	fmt.Println(anotherSum)

}

// Example of a variadic function
// Takes an arbitrary number of arguments of type int
func sumup(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}
