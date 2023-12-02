package main

import "fmt"

func main() {
	greet("HuXn ", "WebDev")
	double(1,2,3)
	ppls := []string{"HuXn", "John", "Jordan"}
	itter(ppls)
}

func greet(firstName string, lastName string) (string, string) {
	return "Hello", firstName + lastName
}

func double(nums ...int) {
	for _, v := range nums {
		fmt.Println(v*2)
	}
}

func itter(slice []string)  {
	for i, v := range slice {
		fmt.Println(i, "-", v)
	}
}