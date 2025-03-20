package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// var price []float64 = []float64{10, 20, 30}
	// objectives - is to calculate the tax for each price
	price := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	priceForEachTaxRate := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceAfterTax := make([]float64, len(price))
		for index, price := range price {
			priceAfterTax[index] = price * (1 + taxRate)
		}
		priceForEachTaxRate[taxRate] = priceAfterTax
	}
	fmt.Println(priceForEachTaxRate, "this is price for each tax rate")
	for key, value := range priceForEachTaxRate {
		for index, priceAfterTax := range value {
			fmt.Println("Tax rate: ", key, "Before tax: ", price[index], "After tax: ", priceAfterTax)
		}
	}
}
