package main

import "fmt"

func main() {
	// Define a map
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}
	fmt.Println(websites)

	// Access the value associated with a key
	fmt.Println(websites["AWS"])

	// Add a key-value pair to the map
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	// Delete a key-value pair from the map
	delete(websites, "Google")
	fmt.Println(websites)

}
