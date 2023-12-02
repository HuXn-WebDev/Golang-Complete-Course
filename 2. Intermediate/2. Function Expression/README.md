# Function Expression

### When we store a function inside a "variable" that function is known as Function Expression.

```go
package main

import "fmt"

func main() {
	add := func(num1 int, num2 int) int {
		return num1 + num2
	}

	res := add(10, 20)
	fmt.Println(res)
}

```
