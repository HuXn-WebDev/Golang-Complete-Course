package main

import "fmt"

func main() {
	// Unbuffered channel
	unbufferedCh := make(chan int)

	// Buffered channel with a capacity of 2
	bufferedCh := make(chan string, 2)

	// Sending data into the unbuffered channel
	go func() {
		unbufferedCh <- 42
	}()

	// Sending data into the buffered channel
	go func() {
		bufferedCh <- "Hello"
		bufferedCh <- "World"
		close(bufferedCh)
	}()

	// Receiving data from the unbuffered channel
	value := <-unbufferedCh
	fmt.Println("Received from unbuffered channel:", value)

	// Receiving data from the buffered channel
	for msg := range bufferedCh {
		fmt.Println("Received from buffered channel:", msg)
	}
}
