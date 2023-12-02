package main

import "fmt"

func main() {
	ppls := []string{"HuXn", "John", "Jordan"}

	for i := 0; i < len(ppls); i++ {
		fmt.Println(ppls[i])
	}

	for index, value := range ppls {
		fmt.Println(index, value)
	}
}
