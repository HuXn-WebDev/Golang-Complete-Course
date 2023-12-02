package main

import "fmt"

func main() {
	// -------------------------------------------
	// 1. Creating Map
	userData := map[string]int{
		"huxn": 17,
		"alex": 18,
		"john": 27,
	}

	// 2. Printing Map
	fmt.Println(userData)

	// 3. Assigning new values to Map
	userData["jordan"] = 15

	// 4. Accessing values from Map
	fmt.Println(userData["huxn"])
	fmt.Println(userData["alex"])
	fmt.Println(userData["john"])

	// 5. Iterating Over Map
	for prop, value := range userData {
		fmt.Println(prop, value)
	}
	fmt.Println(userData)

	// 6. Deleting Items from Map
	delete(userData, "john")
	fmt.Println(userData)
	// -------------------------------------------

	// -------------------------------------------
	studentsMarks := map[string]int{
		"huxn":  30,
		"john":  50,
		"alex":  70,
		"Abdul": 100,
	}

	fmt.Println(studentsMarks)
	fmt.Print(studentsMarks["Abdul"])
	delete(studentsMarks, "Abdul")
	fmt.Println("------------------------")
	fmt.Println(studentsMarks)
	// -------------------------------------------

	// -------------------------------------------
	newStudentsData := map[int]string{
		1: "alex",
		2: "John",
		3: "Jordan",
		4: "HuXn",
		5: "Abdul",
	}

	fmt.Println(newStudentsData[4])
	// -------------------------------------------
}
