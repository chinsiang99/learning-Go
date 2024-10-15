package main

import "fmt"

func main() {
	// websites := []string{"https://google.com", "https://aws.com"}
	// websites := map[string]string // meaning key is string, value is string
	websites := map[string]string{
		"Google":             "https://google.com",
		"Amazon Web Service": "https://aws.com",
	} // meaning key is string, value is string
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	fmt.Println(websites["Amazon Web Service"])
	// to add a new one:
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	// delete a key value pair
	delete(websites, "Google")
	fmt.Println(websites)

	// we have a question here, why maps? why struct?
	// maps offer more flexibility
	// struct has predefined structure, while maps don't
	// struct is to define entities

	// if we don't use make, it will always create a new array and add the values to it if we use append, therefore, we need to tell go that we will need a bigger one, which utilizing make
	userNames := make([]string, 2, 3) // here we define it as length 2, we define the capacity as 3, which is reserving memory, so it will be more efficient, but we can actually go beyond that, but need to have go back to the old days

	userNames = append(userNames, "Max") // here it will still make it further
	userNames = append(userNames, "Christie")

	fmt.Println(userNames)
	fmt.Println(len(userNames))

	// but if we assign directly also can
	userNames[0] = "chin siang"
	fmt.Println(userNames)

	// here it is different with slices, we cant define empty slots, so only one argument
	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.5
	courseRatings["angular"] = 4.5
	fmt.Println(courseRatings)

	courseRatings2 := make(floatMap, 3)
	courseRatings2["go"] = 5.55

	courseRatings2.output()
}

// custom type aliases (better development experience)
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}
