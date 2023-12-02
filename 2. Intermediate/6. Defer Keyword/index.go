package main

import "fmt"

func greet() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}

func main() {
 greet()
}

