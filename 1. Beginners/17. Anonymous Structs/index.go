package main

import "fmt"

func main() {
	h := struct {
		firstName string
		lastName  string
	}{
		firstName: "HuXn",
		lastName:  "WebDev",
	}

	fmt.Println(h.firstName, h.lastName)
}
