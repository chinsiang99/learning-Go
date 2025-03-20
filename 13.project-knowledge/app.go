package main

import (
	"fmt"

	filemanager "chinsiang.com/price-calculator/file-manager"
	"chinsiang.com/price-calculator/prices"
)

func main() {
	fmt.Println("Hello, World!")
	// var price []float64 = []float64{10, 20, 30}
	// objectives - is to calculate the tax for each price
	// price := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	// priceForEachTaxRate := make(map[float64][]float64)

	// for _, taxRate := range taxRates {
	// 	priceAfterTax := make([]float64, len(price))
	// 	for index, price := range price {
	// 		priceAfterTax[index] = price * (1 + taxRate)
	// 	}
	// 	priceForEachTaxRate[taxRate] = priceAfterTax
	// }
	for _, taxRate := range taxRates {
		fileManager := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		job.Process()
		fmt.Println(job.TaxIncludedPrices, "this is price for each tax rate")
	}
	// fmt.Println(priceForEachTaxRate, "this is price for each tax rate")
	// for key, value := range priceForEachTaxRate {
	// 	for index, priceAfterTax := range value {
	// 		fmt.Println("Tax rate: ", key, "Before tax: ", price[index], "After tax: ", priceAfterTax)
	// 	}
	// }
}
