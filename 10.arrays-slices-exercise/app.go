package main

import "fmt"

func main() {
	// 1)
	hobbies := [3]string{"volleyball", "basketball", "badminton"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3)
	// mainHobbies := hobbies[:2]
	mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)

	// 4)
	fmt.Println(cap(mainHobbies))
	// mainHobbies = mainHobbies[1:] // this wont include the last one, because it is according to the existing slice, if we want, then we need to specify it, such as below with 3
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5)
	courseGoals := []string{"rich", "healthy"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "happy"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "fun")
	fmt.Println(courseGoals)

	// 7)
	type Product struct {
		id    int
		title string
		price float64
	}

	dynamicProducts := []Product{Product{id: 1, title: "product 1", price: 100}, Product{id: 2, title: "product 2", price: 200}}
	fmt.Println(dynamicProducts)

	dynamicProducts = append(dynamicProducts, Product{id: 3, title: "product 3", price: 500})

	fmt.Println(dynamicProducts)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
