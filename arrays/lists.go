package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	// Slice with dynamic length
	// In Go, it's common to work with slices right from the start,
	// rather than using arrays.
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	// Append to prices, which create a new slice
	// By reassigning the value to prices, we overwrite it
	prices = append(prices, 5.99)
	fmt.Println(prices)

	// Remove elements by taking slices
	prices = prices[1:]
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(prices[0])

// 	// Slices
// 	// Slices don't create copies of data. It only references pieces of an existing array in memory
// 	featuredPrices := prices[:3]
// 	fmt.Println(featuredPrices)

// 	// Length vs. capacity
// 	// Length: How many elements are in the slice
// 	// Capacity: How many elements are in the parent slice/array that could be selected
// 	//           EXCLUDING any elements to the left of the starting element of the slice
// 	// I.e., you can always select more items to the right, but never to the left
// 	fmt.Println("len(featuredPrices): ", len(featuredPrices))
// 	fmt.Println("cap(featuredPrices): ", cap(featuredPrices))

// }
