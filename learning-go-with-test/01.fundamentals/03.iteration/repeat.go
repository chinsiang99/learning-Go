package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, count int) string {
	// var repeated string
	// for i := 0; i < count; i++ {
	// 	repeated += character
	// }
	// return repeated

	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
