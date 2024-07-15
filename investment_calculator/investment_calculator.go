package main

import (
	"fmt"
	"math"
)

// Declare constants
const inflationRate = 2.5

func main() {
	// Declare variables
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// Get input from user
	outputText("Enter the initial investment amount: $")
	fmt.Scan(&investmentAmount)

	outputText("Enter the number of years: ")
	fmt.Scan(&years)

	outputText("Enter the expected annual rate of return: ")
	fmt.Scan(&expectedReturnRate)

	// Calculate future value, with and without accounting for inflation
	fv, rfv := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// Output the results
	formattedFV := fmt.Sprintf("Future Value: $%.2f\n", fv)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for inflation): $%.2f\n", rfv)

	fmt.Print(formattedFV, formattedFRV)

	// fmt.Printf(`Future Value: $%.2f\n
	// Future Value (adjusted for inflation): $%.2f\n`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(p, r, t float64) (fv float64, rfv float64) {
	fv = p * math.Pow(1+r/100, t)
	rfv = fv / math.Pow(1+inflationRate/100, t)
	return fv, rfv
}
