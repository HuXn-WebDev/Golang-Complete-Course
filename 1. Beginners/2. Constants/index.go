package main

import "fmt"

const (
	a = 10
	b = 20
	c = 30
	d = 40
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var fruit string = "Apple"
	fruit = "Mango" // it will work
	fmt.Println(fruit)

	// You can also write it without specifying type.
	// const apple = "Apple"
	// apple = "Mango" // Error
	// fmt.Println(apple)

	const (
		one   = "John"
		two   = "Alex"
		three = "HuXn"
	)

	// ERROR
	// one = "Alex"
	// two = "John"
	// three = "Nina"

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)

}
