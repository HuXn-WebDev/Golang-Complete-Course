package main

import (
	"fmt"
	"strings"
)

func main() {
	// ---------------------------
	// 1. Using basic string.
	name := "HuXn WebDev"
	// ---------------------------

	// ---------------------------
	// 2. Using Backticks (multiline string).
	intro := `Hello
	  my name is
	  HuXn`
	// ---------------------------

	// ---------------------------
	// 3. String Concatenation.
	firstName := "HuXn "
	lastName := "WebDev"
	// ---------------------------

	// ---------------------------
	// 4. Concatenation In one line
	greetings := "Hello" + " my name is " + "huxn"
	// ---------------------------

	// ---------------------------
	fmt.Println(name)
	fmt.Println(intro)
	fmt.Println(firstName + lastName)
	fmt.Println(greetings)
	// ---------------------------

	// ---------------------------
	// 5. Accessing first character from string.
	message := "HuXn WebDev"
	fmt.Printf("%c\n", message[0])

	message2 := "HuXn WebDev"
	fmt.Printf("%c\n", message2[2])

	lastMessage := "HuXn WebDev"
	fmt.Printf("%c\n", lastMessage[9])
	// ---------------------------

	// ---------------------------
	// 6. Checkout the length of the string
	fruit := "Mango"
	fmt.Println(len(fruit))

	// 7. Comparing two strings
	msg := "one"
	msg2 := "one"
	fmt.Println(strings.Compare(msg, msg2))
	// -1 because string1 is smaller than string2
	// 1 because string2 is greater than string3
	// 0 because string1 and string3 are equal
	// ---------------------------

	// ---------------------------
	// 7. Checking if the string contains certain "character(s)"
	findChar := "Golang Programming Language"
	fmt.Println(strings.Contains(findChar, "Golang")) // true
	fmt.Println(strings.Contains(findChar, "huxn"))   // false
	// ---------------------------

	// 8. Change the text to lowercase
	stringOne := "Hello"
	// stringTwo := "bye"
	fmt.Println(strings.ToLower(stringOne))

	// 9. Change the text to UPPERCASE
	fmt.Println(strings.ToUpper(stringOne))

	// RULES
	// 1. You cannot use single quotes unlike other programming languages.
	// 2. You cannot write your string in multiline using double quotes string, If you still wanna use mutliline string for that you'll have to use backticks `🤚` which will be underneath your escape key and above your tabe key in your keyboard.

}
