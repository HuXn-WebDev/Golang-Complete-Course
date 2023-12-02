# Functions

### A function is a block of statements that can be used repeatedly in a program.

### Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

```go
package main
import "fmt"

// My Own Function
func printName(username string) {
    fmt.Println(username)
}

func main() {
    printName("HuXn")
}
```
