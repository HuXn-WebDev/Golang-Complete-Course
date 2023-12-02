package main

import "fmt"

func checkPassword(password string) {
	if len(password) < 8 {
		panic("Password should be 8 characters")
	} else {
		fmt.Println("Valid Password :)")
	}
}

func main() {
	// checkPassword("12345678") // Valid Password
	checkPassword("125as") // Panic :(
}
