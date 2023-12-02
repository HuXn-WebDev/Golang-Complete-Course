package main

import "fmt"

// Animal interface represents an animal that can provide information
type Animal interface {
	GetInfo() string
}

// Cat struct represents a cat
type Cat struct {
	Name  string
	Color string
}

// Implement the Animal interface for Cat
func (c Cat) GetInfo() string {
	return fmt.Sprintf("Cat: %s, Color: %s", c.Name, c.Color)
}

// Dog struct represents a dog
type Dog struct {
	Name  string
	Breed string
}

// Implement the Animal interface for Dog
func (d Dog) GetInfo() string {
	return fmt.Sprintf("Dog: %s, Breed: %s", d.Name, d.Breed)
}

// Function to print information about an animal
func printAnimalInfo(animal Animal) {
	fmt.Println(animal.GetInfo())
}

func main() {
	// Create instances of Cat and Dog
	cat := Cat{Name: "Whiskers", Color: "Gray"}
	dog := Dog{Name: "Buddy", Breed: "Labrador"}

	// Print information about Cat using the interface
	printAnimalInfo(cat)

	// Print information about Dog using the interface
	printAnimalInfo(dog)
}

