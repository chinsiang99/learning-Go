package main // we declare it as same as main so that inside main, we can call it directly because it is inside the same package

import "fmt"

func presentOptions() {
	fmt.Println("Welcome to Go bank application")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withraw money")
	fmt.Println("4. Exit")
}
