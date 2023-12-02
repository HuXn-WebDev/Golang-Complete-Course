package main

import "fmt"

func addName(name string, callback func(string)) {
	callback(name)
}

func main() {
	addName("HuXn", func(nm string) {
		fmt.Printf("Hi may name is %v \n", nm)
	})
}
