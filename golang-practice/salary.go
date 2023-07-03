package main

import "fmt"

// Struct and methods.

type User struct {
	Name, Email, Role string
	Age               int
}

func (u User) Salary() int {
	switch u.Role {
	case "Developer":
		return 100
	case "Customer Support":
		return 50
	default:
		return 0
	}
}

func (u *User) updateEmail(email string) {
	u.Email = email
}
func main() {
	user := User{
		Name: "Manas",
		Role: "Customer Support",
		Age:  23}

	user.updateEmail("manas@dgraph.io")
	fmt.Println(user)
}
