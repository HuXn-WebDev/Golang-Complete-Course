package main

import "fmt"

func main() {
	usersInfo := map[string]string {
		"huxn": "Full-Stack Developer",
		"john": "Front-End Developer",
		"jordan": "Back-End Developer",
	}

	usersInfo["Kumar"] = "UX/UI Developer"
	fmt.Println(usersInfo)

	delete(usersInfo, "huxn")
	fmt.Println(usersInfo)

	randNums := map[int]int {
		0: 10,
		1: 20,
		2: 30,
		3: 40,
		4: 50,
	}
	
	fmt.Println(randNums)
}