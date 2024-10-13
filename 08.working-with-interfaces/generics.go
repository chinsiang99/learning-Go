package main

func generics() {
	result := add(1, 2)

	result += 1
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
