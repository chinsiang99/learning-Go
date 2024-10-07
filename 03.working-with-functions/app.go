package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("taxRate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	storeResults(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("value must be positive number")
	}
	return userInput, nil
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses         // earning before tax
	profit := ebt * (1 - taxRate/100) // earning after tax
	ratio := ebt / profit

	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
