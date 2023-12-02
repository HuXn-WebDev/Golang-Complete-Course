package main

import "fmt"

func main() {
	num := []int{10, 20, 30, 40, 50}
	fmt.Println(num)

	fmt.Println(num[:])
	fmt.Println(num[1:])
	fmt.Println(num[1:3])

	num = append(num, 60, 70, 80, 90)
	fmt.Println(num)

	num2 := make([]int, 5, 20)
	fmt.Println(num2)

	num2[2] = 5
	fmt.Println(num2)
	// num2[6] = 30
	// fmt.Println(num2)
}
