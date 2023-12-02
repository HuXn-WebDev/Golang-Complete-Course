# Methods

### A method is a function associated with a particular type. It is a way to define behavior for a specific type. In Go, methods are declared with a receiver, which is a special type of parameter that appears in the method signature. The receiver indicates on which type the method operates.

```go
package main

import "fmt"

// Define a type named "Person"
type Person struct {
    FirstName string
    LastName  string
}

// Define a method named "FullName" for the type "Person"
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}

func main() {
    // Create a Person instance
    person := Person{FirstName: "John", LastName: "Doe"}

    // Call the method on the Person instance
    fullName := person.FullName()

    // Print the result
    fmt.Println("Full Name:", fullName)
}


```
