package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from Channel 1"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- "Message from Channel 2"
	}()

	// Use select to wait for either ch1 or ch2
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from Channel 1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from Channel 2:", msg2)
	case <-time.After(5 * time.Second):
		fmt.Println("No communication ready")
	}
}
