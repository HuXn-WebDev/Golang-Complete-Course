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

// --------------------------
// Example 2

// package main

// import "fmt"

// func main() {
// 	address := 10
// 	fmt.Println(address)
// 	fmt.Println(&address)

// 	fmt.Printf("%T\n", address)
// 	fmt.Printf("%T\n", &address)

// 	value := &address
// 	fmt.Println(*&value)
// }