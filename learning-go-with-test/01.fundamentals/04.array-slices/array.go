package arrayslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// There's a new way to create a slice. make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through. The length of a slice is the number of elements it holds len(mySlice), while the capacity is the number of elements it can hold in the underlying array cap(mySlice), e.g., make([]int, 0, 5) creates a slice with length 0 and capacity 5.
// | Use case                                          | Use     |
// | ------------------------------------------------- | ------- |
// | Need actual current elements                      | `len()` |
// | Need total allocated space (e.g., when appending) | `cap()` |

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
