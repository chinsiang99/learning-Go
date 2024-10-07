package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	// var revenue int
	// var expenses int
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scanln(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scanln(&expenses)

	fmt.Print("taxRate: ")
	fmt.Scanln(&taxRate)

	ebt := revenue - expenses         // earning before tax
	profit := ebt * (1 - taxRate/100) // earning after tax
	ratio := ebt / profit

	formattedStringFutureValue := fmt.Sprintf("Future value: %.1f\n", revenue)
	fmt.Print(formattedStringFutureValue)

	fmt.Printf("Future value: %v\n", revenue)
	fmt.Println("ebt: ", ebt)
	fmt.Println("profit: ", profit)
	fmt.Println("ratio: ", ratio)
}
