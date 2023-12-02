package main

import "log"

// Interface
type Human interface {
	// Method Structure
	UserName() string
	Profession() string
}

// Structs
type User1 struct {
	Age int
	Gender string
}

type User2 struct {
	Gender string
	Age int
	Location string
}


// Methods
func (j User1) Profession() string {
	return "Front-End Developer"
}

func (j User1) UserName() string {
	return "John"
}

func (a User2) Profession() string {
	return "Back-End Developer"
}

func (a User2) UserName() string {
	return "Lara"
}

func main() {
	// Instance of structs
	john := User1 {
		Gender: "male",
		Age: 20,
	}

	PrintInfo(john)

	Lara := User2 {
		Gender: "female",
		Age: 19,
		Location: "USA",
	}

	PrintInfo(Lara)
}

// Info Function
func PrintInfo(h Human) {
	log.Println("My name is a ", h.UserName(), "and i am a professional", h.Profession())
}
