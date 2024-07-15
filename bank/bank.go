package main

import "fmt"

// Define global balance variable
var balance float64 = 0

func main() {
	printHeader()

	// Loop indefinitely until user exits
	for {
		showPrompt()
		choice := getUserChoice()
		switch choice {
		case 1:
			checkBalance()
		case 2:
			deposit(getUserAmount("Amount to deposit: "))
		case 3:
			withdraw(getUserAmount("Amount to withdraw: "))
		default:
			printGoodbye()
			return
		}
	}
}

// Print header once at beginning of execution
func printHeader() {
	fmt.Println("Welcome to Johnstone Bank!")
	fmt.Println("--------------------------")
}

// Print current balance
func checkBalance() {
	fmt.Printf("\nYour current balance is: $%.2f\n", balance)
}

// Withdraw a given amount from balance, if able
func withdraw(amount float64) {
	fmt.Println()
	if amount < 0 {
		fmt.Println("Amount to withdraw must be greater than 0!")
	} else if amount > balance {
		fmt.Println("Not enough money in your account!")
	} else {
		balance -= amount
		fmt.Printf("Withdrawing $%.2f, balance is now $%.2f.\n", amount, balance)
	}
}

// Deposit given amount to balance
func deposit(amount float64) {
	fmt.Println()
	if amount < 0 {
		fmt.Println("Amount to deposit must be greater than 0!")
	} else {
		balance += amount
		fmt.Printf("\nDepositing $%.2f, balance is now $%.2f.\n", amount, balance)
	}
}

// Get option choice from the user
func getUserChoice() (choice int) {
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return
}

// Get a numerical amount from user
func getUserAmount(prompt string) (amount float64) {
	fmt.Print(prompt)
	fmt.Scan(&amount)
	return
}

// Print user prompt, once per iteration of the main loop
func showPrompt() {
	fmt.Println("\nWhat would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println()
}

// Print exit text
func printGoodbye() {
	fmt.Println("Thank you for banking with us. Goodbye!")
}
