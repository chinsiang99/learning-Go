package main

import (
	"fmt"
	"os"
)

func main() {
	var accountBalance float64 = 1000

	// for in Go is the only loop you can use
	// for i := 0; i < 2; i++ {
	for {
		fmt.Println("Welcome to Go bank application")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withraw money")
		fmt.Println("4. Exit")

		fmt.Print("Please select your choice: ")
		var choice int
		fmt.Scan(&choice)

		fmt.Println("Your choice is:", choice)

		// wantsCheckBalance := choice == 1

		if choice == 1 {
			fmt.Println("Your account balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			// accountBalance = accountBalance + depositAmount
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				// return // can be treated as stopping the function here
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Print("Your Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			// accountBalance = accountBalance - withdrawalAmount
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				// return // can be treated as stopping the function here
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("You canot withdraw amount greater than account balance")
				// return // can be treated as stopping the function here
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else {
			fmt.Println("Thanks for using Go bank application, have a nice day!")
			// return // if we are using return instead, it will stop the function and will not execute next steps, therefore, break is been used here
			break
		}
	}

	fmt.Println("Thanks for choosing our bank application!")
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644) // 0644 is file permission
}
