## Go Conditions

### Conditional statements allow us to control the structure of our program.

### There are different ways by which we can control the flow of our program, (If, else if, else) are one of them.

### (If, else if, else) statments allow us to make "decisions" while our program is running, They're also called (conditional statments) in programming.

```go
    // Sudo Syntax
     if condition { <code> }
     else if condition { <code> }
     else { <code> }
```

```go
// Actual Code
package main
import "fmt"

func main() {
	password := "12345678"
	if len(password) > 7 {
		fmt.Println("Valid Password")
	} else {
		fmt.Println("Invalid Password")
	}
}
```
