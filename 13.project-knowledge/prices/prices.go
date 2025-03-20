package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("Before tax: %.2f After tax: %.2f", price, price*(1+job.TaxRate))] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = result

	// fmt.Println(job.TaxIncludedPrices)
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	errScanner := scanner.Err()

	if errScanner != nil {
		fmt.Println("Error reading file", err)
		file.Close()
		return
	}
}
