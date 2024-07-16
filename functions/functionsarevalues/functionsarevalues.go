package functionsarevalues

import "fmt"

// Use custom type to alias function types
type transformFn func(int) int

// Example alias that would save a of typing
// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	// Pass a function as a parameter
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	// Return a function from a function...
	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	// ..then pass those functions into another function
	transformedNums1 := transformNumbers(&numbers, transformerFn1)
	transformedNums2 := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNums1)
	fmt.Println(transformedNums2)
}

// Define a function that takes another function as an argument
// Must specify the expected inputs and outputs
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

// Function that returns another function
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
