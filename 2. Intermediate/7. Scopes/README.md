# Scope

### Think of scope as the visibility or accessibility of things (like variables or functions) in your code. It's like where things are known or can be used.

### 1. Block Scope

#### Imagine your code as a series of blocks, like paragraphs in a story.

#### Variables and functions created inside a block are known only within that block.

```go
func main() {
    // Inside main's block
    var x int = 10
    fmt.Println(x)  // Can use x here

    if true {
        // Inside if's block
        var y int = 20
        fmt.Println(y)  // Can use y here
    }

    // Can't use y here, it's outside its block
}
```

### 2. Function Scope

#### Functions have their own scope. Variables defined inside a function are known only within that function.

```go
func myFunction() {
    // Inside myFunction's block
    var z int = 30
    fmt.Println(z)  // Can use z here
}

// Can't use z here, it's outside myFunction's block
```

### 3. Package Scope

#### Variables and functions defined at the top level of a package have a wider scope.

#### They can be used across multiple files within the same package.

```go
package main

// Outside any function, package scope
var globalVariable int = 50

func main() {
    // Can use globalVariable here
}
```
