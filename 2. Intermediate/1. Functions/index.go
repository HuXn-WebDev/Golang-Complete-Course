package main

import "fmt"

func main() {
	hello()
	greet("John")
	showNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	showUsers("HuXn", "Alex", "John", "Jordan")
}

func hello() {
	fmt.Println("Hello")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func showNumbers(numbers ...int) {
	fmt.Println(numbers)
}

func showUsers(s ...string) {
	for i, v := range s {
		fmt.Println(i, "-", v)
	}
}
