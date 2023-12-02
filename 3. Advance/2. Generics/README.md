# Generics

### Generics in Go provide a way to write functions and data structures that can work with any data type while maintaining type safety

### 1. Generics in Go involve defining functions or data structures with type parameters.

### 2. Type parameters are specified inside square brackets ([]) before the function or type signature.

```go
package main

import "fmt"

// PrintSlice is a generic function that prints elements of any slice.
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {
    // Example with integers
    intSlice := []int{1, 2, 3, 4, 5}
    PrintSlice(intSlice)

    // Example with strings
    stringSlice := []string{"apple", "banana", "cherry"}
    PrintSlice(stringSlice)
}
```
