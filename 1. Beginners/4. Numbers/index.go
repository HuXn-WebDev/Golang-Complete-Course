package main

import "fmt"

func main() {
	// var num1 int = 20000000
	var num1 int32 = 20000000
	fmt.Println(num1)

	// var num2 float32 = 4545654.5555
	var num2 float64 = 4545654.5555
	fmt.Println(num2)

	// You can also ommit the types and
	// That variable will by default "inferred" it.
	specialNumber := 74353.0
	fmt.Printf("%T", specialNumber) // checking the type of variable.

	// Arithmetic Operators 👇
	fmt.Println(2 + 2)
	fmt.Println(2 - 2)
	fmt.Println(2 * 2)
	fmt.Println(2 / 2)
	fmt.Println(2 % 2)
}
