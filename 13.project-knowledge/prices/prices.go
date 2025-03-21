package prices

import (
	"fmt"
	"math"

	"chinsiang.com/price-calculator/conversion"
	filemanager "chinsiang.com/price-calculator/file-manager"
)

type TaxIncludedPriceJob struct {
	IOManager         *filemanager.FileManager `json:"-"`
	TaxRate           float64                  `json:"tax_rate"`
	InputPrices       []float64                `json:"input_prices"`
	TaxIncludedPrices map[string]float64       `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(fm *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("Before tax: %.2f After tax: %.2f", price, price*(1+job.TaxRate))] = math.Round(price * (1 + job.TaxRate))
	}

	job.TaxIncludedPrices = result

	// filemanager.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	job.IOManager.WriteJson(job)

	// fmt.Println(job.TaxIncludedPrices)
}

func (job *TaxIncludedPriceJob) LoadData() {
	// file, err := os.Open("prices.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file", err)
	// 	return
	// }
	// defer file.Close()

	// lines := []string{}
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }

	// errScanner := scanner.Err()

	// if errScanner != nil {
	// 	fmt.Println("Error reading file", err)
	// 	file.Close()
	// 	return
	// }

	// lines, err := filemanager.ReadLines("prices.txt")
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	// prices := make([]float64, len(lines))
	// for lineIndex, line := range lines {
	// 	floatPrice, err := strconv.ParseFloat(line, 64)
	// 	if err != nil {
	// 		fmt.Println("Error parsing price to float64", err)
	// 		file.Close()
	// 		return
	// 	}
	// 	prices[lineIndex] = floatPrice
	// }

	// prices, err := conversion.StringsToFloat64(lines)
	prices, err := conversion.ToFloat64Slice(lines)
	if err != nil {
		fmt.Println("Error converting strings to float64", err)
		return
	}

	job.InputPrices = prices
}
