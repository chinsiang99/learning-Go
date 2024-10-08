package user

import (
	"errors"
	"fmt"
	"time"
)

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// type User // allow export out of this package by defining uppercase first letter

type User struct {
	firstName string
	// FirstName string // this FirstName instead of firstName because it wanted to be accessed outside (being export)
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

// pointer here by making sure that it will change the original struct
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName string, lastName string, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
	}, nil
}
