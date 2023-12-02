# Interface

### An interface is a type that specifies a set of method signatures. It defines a contract for what methods a type must have, but it does not provide the implementation for those methods. Instead, the implementation is provided by concrete types that satisfy the interface.

```go

package main

import "fmt"

// Define an interface named Speaker
type Speaker interface {
	Speak() string
}

// Implement the Speaker interface for Dog
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Implement the Speaker interface for Cat
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Function that takes any type implementing the Speaker interface
func announce(speaker Speaker) {
	fmt.Println("The speaker says:", speaker.Speak())
}

func main() {
	// Create instances of Dog and Cat
	dog := Dog{}
	cat := Cat{}

	// Use the announce function with different types
	announce(dog) // The speaker says: Woof!
	announce(cat) // The speaker says: Meow!
}


```
