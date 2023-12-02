# The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

```go

package main
import ("fmt")

const user = "admin" // cannot be changed

func main() {
  fmt.Println("admin")
}

```

# Constant Rules

### 1. Constant names follow the same naming rules as variables

### 2. Constant names are usually written in uppercase letters

### 3. Constants can be declared both inside and outside of a function
