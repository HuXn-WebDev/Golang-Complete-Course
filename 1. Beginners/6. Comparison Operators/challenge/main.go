package main

import "fmt"

var number = 20

func main() {
	fmt.Println(number > 10)
	fmt.Println(number < 10)
	fmt.Println(number >= 10)
	fmt.Println(number <= 10)
	fmt.Println(number == number)
	fmt.Println(number != number)
}