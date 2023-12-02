package main

import "fmt"

type human struct {
	gender   string
	fullName string
	age      int
	friends
}

type friends struct {
	justFriends int
	bestFriend  string
}

// Method
func (h human) printInfo() {
	fmt.Println("My name is", h.fullName, "Gender:", h.gender, "Age:", h.age, "I have", h.justFriends, "Friends but my best friend is", h.bestFriend)
}

func main() {
	huxn := human{
		gender:   "male",
		fullName: "HuXn WebDev",
		age:      17,
		friends: friends{
			justFriends: 20,
			bestFriend:  "John",
		},
	}
	huxn.printInfo()
}
