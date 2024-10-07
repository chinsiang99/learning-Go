package main

import (
	"fmt"

	"chinisang.com/simple-bank-app-split/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	// var accountBalance float64 = 1000
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-------------------")
		// panic("Can't continue, sorry")
	}

	// for in Go is the only loop you can use
	// for i := 0; i < 2; i++ {
	for {
		presentOptions()

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
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
		} else {
			fmt.Println("Thanks for using Go bank application, have a nice day!")
			// return // if we are using return instead, it will stop the function and will not execute next steps, therefore, break is been used here
			break
		}
	}

	fmt.Println("Thanks for choosing our bank application!")
}

// func writeBalanceToFile(value float64, fileName string) {
// 	valueText := fmt.Sprint(value)
// 	os.WriteFile(fileName, []byte(valueText), 0644) // 0644 is file permission
// }

// func getFloatFromFile(fileName string) (float64, error) {
// 	data, err := os.ReadFile(fileName)

// 	if err != nil {
// 		return 1000, errors.New("failed to find file")
// 	}

// 	valueText := string(data)
// 	// float64(valueText)
// 	value, err := strconv.ParseFloat(valueText, 64)
// 	if err != nil {
// 		return 1000, errors.New(("failed to parse stored value"))
// 	}
// 	return value, nil
// }

// func presentOptions() {
// 	fmt.Println("Welcome to Go bank application")
// 	fmt.Println("What do you want to do?")
// 	fmt.Println("1. Check balance")
// 	fmt.Println("2. Deposit money")
// 	fmt.Println("3. Withraw money")
// 	fmt.Println("4. Exit")
// }
