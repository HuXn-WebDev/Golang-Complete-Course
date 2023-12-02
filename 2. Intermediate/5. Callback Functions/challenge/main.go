package main

import (
	"fmt"
)

func main() {
    run(hello)
    run(world)
}

func hello() {
    fmt.Println("Hello")
}

func world() {
    fmt.Println("World!")
}

func run(f func()) {
    f()
}