package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// note that if our package is lower case, we are not able to export the function
func WriteBalanceToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644) // 0644 is file permission
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	// float64(valueText)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New(("failed to parse stored value"))
	}
	return value, nil
}
