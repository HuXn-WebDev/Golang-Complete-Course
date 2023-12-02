package main

import "fmt"

func multiply(num1, num2 int) {
    fmt.Println("Result: ", num1 * num2)
}
 
func printMessage() {
    fmt.Println("Hello, HuXn")
}
 
func main() {
    defer multiply(23, 56)
    printMessage()
}