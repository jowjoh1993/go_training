package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)
	quad := createTransformer(4)

	// Example of anonymous function definition
	transformed := transformNumbers(&numbers, func(n int) int {
		return n * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	quadded := transformNumbers(&numbers, quad)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(quadded)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Example of function closure
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
