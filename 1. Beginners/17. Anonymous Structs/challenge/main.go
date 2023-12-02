package main

import "fmt"

func main() {
	type comments struct {
		author string
		body string
	}

	m := struct {
		title string
		mType string
		rating int
		comments
	} {
		title: "The Social Network",
		mType: "Programming",
		rating: 20,
		comments: comments {
			author: "HuXn",
			body: "Awesome Movie",
		},
	}

	fmt.Println("My Favorite Movie " + "(" + m.title + ")" + " has", m.rating, "Ratings")
	fmt.Println(m.comments.author + " say's : " + m.comments.body)
}