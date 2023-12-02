package main

import "fmt"

func main() {
	role := "guest"
	switch {
	case (role == "guest"):
		fmt.Println("Guest User")
	case (role == "moderator"):
		fmt.Println("Moderator User")
	default:
		fmt.Println("Unknown User")
	}
}
