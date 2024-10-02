package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello there~ the program started hehe")
	const inflationRate float64 = 2.5
	// var investmentAmount float64 = 1000
	var investmentAmount float64
	// var expectedReturnRate float64 = 5.5
	expectedReturnRate := 3.0 // infer type recommendation shorter way assigning a variable
	var years float64 = 10

	// var investmentAmount, years float64 = 1000, 10
	// investmentAmount, years := 1000.0, 10.0
	fmt.Print("Please input your investment amount: ")
	fmt.Scan(&investmentAmount) // & meaning pointers

	fmt.Print("Please input your expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Please input your year(s) of investment: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years) // can convert from int to float64 (use float64(years))

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

// func precision(a) {
// 	if (!isFinite(a)) return 0;
// 	var e = 1, p = 0;
// 	while (Math.round(a * e) / e !== a) { e *= 10; p++; }
// 	return p;
//   }
