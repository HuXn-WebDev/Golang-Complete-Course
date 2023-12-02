package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type programmer struct {
	person
	isProgrammer bool
}

func main() {
	h := programmer{
		person: person{
			firstName: "HuXn",
			lastName:  "WebDev",
		},
		isProgrammer: true,
	}

	fmt.Println(h.firstName, h.lastName, h.isProgrammer)
}
