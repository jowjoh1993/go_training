package main

import "fmt"

// Time to practice what you learned!

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Juggling", "Piano", "Games"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	hobbies1 := hobbies[:2]
	hobbies2 := hobbies[0:2]
	fmt.Println(hobbies1, hobbies2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	hobbies3 := hobbies2[1:3]
	fmt.Println(hobbies3)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Golang proficiency", "REST API experience"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Knowledge of structs"
	goals = append(goals, "More opportunities at work")

	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type Product struct {
		title string
		id    string
		price float64
	}

	products := []Product{
		{"Book", "01234", 12.99},
		{"Hat", "56789", 24.99},
	}

	products = append(products, Product{"Keyboard", "24680", 76.99})

	fmt.Println(products)
}
