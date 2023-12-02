# Comparison Operators

## Comparison operators are used to compare two values

| Operator | Name                     | Example |
| -------- | ------------------------ | ------- |
| ==       | Equal to                 | x == y  |
| !=       | Not equal                | x != y  |
| >        | Greater than             | x > y   |
| <        | Less than                | x < y   |
| >=       | Greater than or equal to | x >= y  |
| <=       | Less than or equal to    | x <= y  |

```go

package main

import "fmt"

func main() {
	fmt.Println(2 > 2) // false
	fmt.Println(2 < 2) // false
	fmt.Println(2 >= 2) // true
	fmt.Println(2 <= 2) // true
	fmt.Println(2 == 2) // true
	fmt.Println(2 != 2) // false
}
```
