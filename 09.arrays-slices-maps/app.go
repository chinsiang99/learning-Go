package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"hello"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	productNames[2] = "product name hehe"
	fmt.Println(productNames)

	fmt.Println(prices[0])

	// first value is included, second value is excluded
	featuredPrices := prices[1:3]
	featuredPrices[0] = 199.99
	fmt.Println(prices)
	// featuredPrices := prices[1:]
	// featuredPrices := prices[:3]
	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	// note that in Go, slices always remember the right available memory of the array, so we can use slice to get an array's right hand side even it only have 1 value (if the original one has more than that)

	fmt.Println("dynamic array started here......")
	// below is working with dynamic array
	dynamicPrices := []float64{10.99, 8}
	fmt.Println(dynamicPrices[1])
	fmt.Println(dynamicPrices[0:1])

	// it will create a brand new array and add that brand new value to the array
	updatedPrices := append(dynamicPrices, 5.99)
	fmt.Println(updatedPrices, dynamicPrices)

	// if we want to amend the original array, we need to reassign it, Go will help us to do the memory management instead
	dynamicPrices = dynamicPrices[1:]
	fmt.Println(dynamicPrices)
}
