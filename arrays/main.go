package main

import "fmt"

// Type alias example: map[string]float64 --> floatMap
type floatMap map[string]float64

// Print method for floatMap custom type
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// Use make() to pre-allocate space in memory for arrays/slices
	userNames := make([]string, 2, 5)

	userNames[0] = "Max"
	userNames[1] = "Manuel"
	userNames = append(userNames, "Julie")

	fmt.Println(userNames)

	// Using make() to pre-allocate memory for maps
	// Only 1 argument for maps, since maps don't have empty elements
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	// Once we add another key-value pair beyond the pre-allocated
	// capacity, Go has to re-allocate memory for this object
	// courseRatings["node"] = 4.6

	// Using for loops to iterate over slices and maps
	for key, value := range courseRatings {
		fmt.Printf("%v: %.1f\n", key, value)
	}

}
