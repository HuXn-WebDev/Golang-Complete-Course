package main

import "fmt"

func main() {
	numbers := []int{42, 21, 11, 23, 44, 56, 20, 100, 200}
	// fmt.Println(numbers)
	fmt.Println(numbers[:])
	fmt.Println(numbers[0:5])
	numbers = append(numbers, 22, 12, 34, 11, 67, 90)
	fmt.Println(numbers)
	mySlice := make([]int, 5, 20)
	mySlice[2] = 20
	fmt.Println(mySlice)
}