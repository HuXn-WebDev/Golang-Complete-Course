package main

import "fmt"

type movie struct {
	title string
	mType string
	rating int
	comments
}

type comments struct {
	author string
	body string
}

func main() {
	m := movie {
		title: "The Social Network",
		mType: "Programming",
		rating: 20,
		comments: comments {
			author: "HuXn",
			body: "Awesome Movie",
		},
	}

	// fmt.Println(m)
	fmt.Println("My Favorite Movie " + "(" + m.title + ")" + " has", m.rating, "Ratings")
	fmt.Println(m.comments.author + " say's : " + m.comments.body)
}