package main

import "fmt"

var name = "HuXn" // Global

func main() {
	m := 200 // Local Scope Variable
	fmt.Println(name)
	showName()
	fmt.Println(m)
}

// Global Scope Function
func showName() {
	fmt.Println(name)
}
