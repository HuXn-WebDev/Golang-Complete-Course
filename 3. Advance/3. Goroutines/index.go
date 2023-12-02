package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	// Launching a goroutine
	go printNumbers()

	// This line will be executed concurrently with the goroutine
	fmt.Println("This is the main function.")

	// Sleeping to allow the goroutine to complete
	time.Sleep(1 * time.Second)
}
