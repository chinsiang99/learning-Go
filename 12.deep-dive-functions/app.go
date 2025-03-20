package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)
	fmt.Println(numbers)
	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := make([]int, len(*numbers)) // Create a new slice to store doubled values
	for i := range *numbers {
		(*numbers)[i] = (*numbers)[i] * 2 // Correctly modify the original slice
		dNumbers[i] = (*numbers)[i] * 2   // Store doubled values in the new slice
	}
	return dNumbers

	//Why Parentheses Are Required:
	//Without parentheses, if you write *numbers[i], Go interprets this as "get the i-th element of the slice pointer numbers, and then dereference that element." But since numbers is a pointer to a slice (*[]int), not a slice of pointers, this is invalid.
}
