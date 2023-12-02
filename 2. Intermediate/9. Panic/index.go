package main

func employee(name *string, age int) {
	if age > 65 {
		panic("Age cannot be greater than retirement age")
	}
}

func main() {
	empName := "John"
	age := 75
	employee(&empName, age)
}