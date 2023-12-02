package main

import "fmt"

func sendData(ch chan<- int) {
	ch <- 42
}

func main() {
	// Creating an unbuffered channel
	ch := make(chan int)

	// Launching a goroutine to send data into the channel
	go sendData(ch)

	// Receiving data from the channel
	value := <-ch

	// Printing the received value
	fmt.Println("Received:", value)
}
