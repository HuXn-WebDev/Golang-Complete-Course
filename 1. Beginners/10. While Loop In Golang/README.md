# While Loop

## Unlike other programming languages, Go doesn't have a dedicated keyword for a while loop. However, we can use the for loop to perform the functionality of a while loop.

```go
// Program to print numbers between 0 and 10
package main
import ("fmt")

func main() {
  number := 0

  for number <= 10 {
    fmt.Println(number)
    number++
  }
}
```
