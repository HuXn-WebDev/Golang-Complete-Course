package main

import "fmt"

func main() {
	double := func(nums ...int) {
		for _, v := range nums {
			fmt.Println(v*2)
		}
	}
	double(1,2,3,4)
}
