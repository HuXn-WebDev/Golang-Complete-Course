package main

import "fmt"

// Define a simple interface named 'Describer'
type Describer interface {
    Describe() string
}

// Implement the interface for a 'Person'
type Person struct {
    Name string
}

// Implement the 'Describe' method for 'Person'
func (p Person) Describe() string {
    return "Hi, I'm " + p.Name
}

// Implement the interface for a 'Car'
type Car struct {
    Model string
}

// Implement the 'Describe' method for 'Car'
func (c Car) Describe() string {
    return "This is a " + c.Model + " car"
}

// Function to display description using the interface
func displayDescription(d Describer) {
    fmt.Println(d.Describe())
}

func main() {
    // Create instances of Person and Car
    person := Person{Name: "Alice"}
    car := Car{Model: "Toyota"}

    // Use the Describe method through the Describer interface
    displayDescription(person)
    displayDescription(car)
}
