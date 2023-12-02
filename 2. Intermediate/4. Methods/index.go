package main

import "fmt"

type person struct {
	gender     string
	firstName  string
	lastName   string
	age        int
	profession string
	isMarried  bool
}

func (p person) printInfo() {
	fmt.Println("I'm", p.firstName, p.lastName, "Gender:", p.gender, "Age:", p.age, "Profession:", p.profession, "Married", p.isMarried)
}

func main() {
	huxn := person{
		gender:     "Male",
		firstName:  "HuXn",
		lastName:   "WebDev",
		age:        17,
		profession: "Programming",
		isMarried:  false,
	}
	huxn.printInfo()
}
