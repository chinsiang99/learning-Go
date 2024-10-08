package main

import (
	"fmt"
	"time"
)

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

// type User // allow export out of this package by defining uppercase first letter

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

// pointer here by making sure that it will change the original struct
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	// empty value
	// appUser = user{}

	// struct literal
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		// createdAt: time.Now(),
	}

	// fmt.Println(firstName, lastName, birthDate)
	// outputUserDetails(&appUser)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}
