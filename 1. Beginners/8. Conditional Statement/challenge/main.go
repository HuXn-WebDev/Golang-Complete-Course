package main

import "fmt"

func main() {
	hour := 12
	if hour >= 6 && hour < 12 {
		fmt.Println("Good Morning")
	} else if hour >= 12 && hour < 18 {
		fmt.Println("Good Afternoon")
	} else {
		fmt.Println("Good Evening")
	}
}
