# Logical Operators

## Logical operators are used to determine the logic between variables or values.

## 1. Logical and (&&)

### Returns true if both statements are true

## 2. Logical or (||)

### Returns true if one statements is true

## 3. Logical not (!)

### Reverse the result, returns false if the result is true

```go

package main

import "fmt"

func main() {
	fmt.Println(true && true) // true
	fmt.Println(true && false) // false

	fmt.Println(true || true) // true
	fmt.Println(false || false) // false

	fmt.Println(!true) // false
	fmt.Println(!false) // true
}
```
