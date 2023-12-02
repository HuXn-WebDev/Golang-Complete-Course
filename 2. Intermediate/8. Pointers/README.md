# Pointers

### A pointer in Go is a variable that stores the memory address of another variable. By using pointers, you can indirectly access and modify the value of the variable whose address is stored in the pointer.

#### You can create a pointer using the \* (asterisk) symbol.

#### The & (ampersand) symbol is used to obtain the memory address of a variable

```go
package main

import "fmt"

func main() {
    // Declare a variable
    x := 42

    // Create a pointer that stores the memory address of x
    var pointerToX *int = &x

    // Print the value and memory address of x
    fmt.Println("Value of x:", x)
    fmt.Println("Memory address of x:", &x)

    // Print the value and memory address stored in the pointer
    fmt.Println("Value pointed to by pointerToX:", *pointerToX)
    fmt.Println("Memory address stored in pointerToX:", pointerToX)
}

```
