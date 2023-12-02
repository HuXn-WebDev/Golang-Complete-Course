package main

import (
	"fmt"
	"sync"
)

func printMessage(wg *sync.WaitGroup, msg string) {
	defer wg.Done() // Decrement the WaitGroup counter when done
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

	wg.Add(3) // Add the number of Goroutines you're waiting for

	go printMessage(&wg, "Hello")
	go printMessage(&wg, "Go")
	go printMessage(&wg, "WaitGroup")

	wg.Wait() // Wait for all Goroutines to finish

	fmt.Println("All messages printed.")
}
