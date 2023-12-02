package main

import "fmt"

func main() {
	// 1. Creating Array
	var numbers [5]int
	fmt.Println(numbers)

	// Adding Items to array
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println(numbers)

	// 3. Creating String array
	var peoples [3]string
	peoples[0] = "HuXn" // Adding items to string array
	peoples[1] = "John" // Adding items to string array
	peoples[2] = "Jordan"
	fmt.Println(peoples)

	// 4. Length & Capcity
	fmt.Println(len(peoples))
	fmt.Println(cap(peoples))

	// 5. Iterating over array.
	for item := 0; item < len(peoples); item++ {
		// fmt.Println(item)
		fmt.Println(peoples[item])
	}
}
