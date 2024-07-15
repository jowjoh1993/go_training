package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1. Validate user input (show error and exit)
//   - No negative numbers
//   - Non zero
// 2. Store calculated results into a file

func main() {
	// Get user input
	revenue, err := userInput("Enter revenue: ")
	if err != nil {
		fmt.Println(err)
		fmt.Println("----------")
		return
	}
	expenses, err := userInput("Enter expenses: ")
	if err != nil {
		fmt.Println(err)
		fmt.Println("----------")
		return
	}
	taxRate, err := userInput("Enter tax rate: ")
	if err != nil {
		fmt.Println(err)
		fmt.Println("----------")
		return
	}

	// Calculate EBT, profit, and ratio
	EBT, profit, ratio := calcMetrics(revenue, expenses, taxRate)

	// Output results to a file
	writeToFile("metrics.txt", EBT, profit, ratio)
}

// Prompts user for input, then passes input to pointer
func userInput(prompt string) (value float64, err error) {
	fmt.Print(prompt)
	fmt.Scan(&value)
	if value < 0 {
		err = errors.New("input cannot be a negative number")
	} else if value == 0 {
		err = errors.New("input cannot be zero")
	} else {
		err = nil
	}
	return value, err
}

// Calculates EBT, profit after tax, and EBT / profit
func calcMetrics(rev, exp, tax float64) (ebt, profit, ratio float64) {
	ebt = rev - exp
	profit = ebt * (1 - tax/100)
	ratio = ebt / profit
	return
}

func writeToFile(filename string, ebt, profit, ratio float64) {
	line1 := fmt.Sprintf("EBT = $%.2f\n", ebt)
	line2 := fmt.Sprintf("Profit = $%.2f\n", profit)
	line3 := fmt.Sprintf("Ratio = %.4f\n", ratio)
	data := fmt.Sprint(line1, line2, line3)
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Results written to %v.\n", filename)
}
