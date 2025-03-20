package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

// func StringsToFloat64(strings []string) ([]float64, error) {
// 	floatResults := make([]float64, len(strings))
// 	for stringIndex, string := range strings {
// 		floatPrice, err := strconv.ParseFloat(string, 64)
// 		if err != nil {
// 			fmt.Println("Error parsing price to float64", err)
// 			return nil, err
// 		}
// 		floatResults[stringIndex] = floatPrice
// 	}
// 	return floatResults, nil
// }

// Generic function to convert a slice of any type (string, int, float64) to float64
func ToFloat64Slice[T any](input []T) ([]float64, error) {
	floatResults := make([]float64, len(input))

	for i, v := range input {
		switch value := any(v).(type) {
		case string:
			floatVal, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Println("Error parsing string to float64:", err)
				return nil, errors.New("error parsing string to float64")
			}
			floatResults[i] = floatVal
		case int:
			floatResults[i] = float64(value)
		case float64:
			floatResults[i] = value
		default:
			return nil, errors.New("Unsupported type of" + fmt.Sprintf("%T", value))
		}
	}
	return floatResults, nil
}
