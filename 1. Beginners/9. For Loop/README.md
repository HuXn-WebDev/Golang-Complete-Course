# For Loop

## A "For" Loop is used to repeat a specific block of code a known number of times.

## For example, if we want to check the grade of every student in the class, we loop from 1 to that number. When the number of times is not known before hand, we use a "While" loop.

```go
// Sudo Syntax
for initialExpression; condition; increment { <code> }
```

```go
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
```

# The continue Statement

## The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

```go
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }
}
```

# The break Statement

## The break statement is used to break/terminate the loop execution.

```go
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }
}
```

# Nested Loops

## Loop inside other loop is known as Nested Loop.

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println("-----OUTER------", i)
		for j := 0; j < 10; j++ {
			fmt.Println("-----INNER------", j)
		}
	}
}

```
