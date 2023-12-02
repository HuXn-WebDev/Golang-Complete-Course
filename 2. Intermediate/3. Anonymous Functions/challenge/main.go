package main

import "fmt"

func main() {
	func(nums ...int) {
		for _, v := range nums {
			fmt.Println(v * 2)
		}
	}(1, 2, 3, 4)
}
