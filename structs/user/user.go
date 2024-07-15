package user

import (
	"errors"
	"fmt"
	"time"
)

// Type name is capitalized so that it's exported from package
// Struct elements are lowercase, and are NOT exported.
// Instead, users must use the struct methods to interact with
// the elements of the struct.
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Create another struct with the User struct embedded within
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

// Print elements of the struct
// Capitalized - we export this method
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Clear the first name and last name
// Capitalized - we export this method
// IMPORTANT: In order for a method to modify the data in a struct,
// it has to receive a pointer to that struct. Otherwise it will
// just get a copy of the object and modify that copy.
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Constructor function, commonly just named "New"
func New(firstName, lastName, birthDate string) (*User, error) {
	// Data validation
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birthDate are required")
	}

	// Build struct
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
