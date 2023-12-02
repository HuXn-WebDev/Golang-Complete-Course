package main

import "fmt"

func main() {
	apple := "apple"
	orange := "orange"

	fmt.Println(&apple)
	fmt.Println(&orange)

	fmt.Println(apple)
	fmt.Println(orange)

	swap := *&apple
	apple = orange
	orange = swap
	fmt.Println("--------------")
	fmt.Println(*&apple)
	fmt.Println(*&orange)
}
