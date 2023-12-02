# Variables are containers for storing data values

## In Go, there are two ways to declare a variable

### 1. Use the var keyword, followed by variable name and type

```go
package main
import ("fmt")

var name = "John Doe" // Variable One

func main() {
  var fruit = "Apple" // Variable Two
  fmt.Println(name)
  fmt.Println(fruit)
}
```

### 2. Use the := sign, followed by the variable value

```go
package main
import ("fmt")

func main() {
  varOne := 100
  varTwo := 2
  fmt.Println(varOne)
  fmt.Println(varTwo)
}
```

### Assigning Types + Type Inferred

```go
package main
import ("fmt")

func main() {
  var fruit string = "Mango" // type is string
  var user = "HuXn" // type is inferred
  number := 2 // type is inferred

  fmt.Println(fruit)
  fmt.Println(user)
  fmt.Println(number)
}
```

### Go Multiple Variable Declaration

```go
package main
import ("fmt")

func main() {
  var one, two, three, four, five int = 1, 2, 3, 4, 5

  fmt.Println(one)
  fmt.Println(two)
  fmt.Println(three)
  fmt.Println(four)
}
```

### Go Variable Declaration in a Block

```go
package main
import ("fmt")

func main() {
   var (
     num int
     number int = 1
     greetings string = "hello"
   )

  fmt.Println(num)
  fmt.Println(number)
  fmt.Println(greetings)
}
```

## Go Variable Naming Rules

#### 1. A variable name must start with a letter or an underscore character (\_)

#### 2. A variable name cannot start with a digit

#### 3. A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and \_ )

#### 4. Variable names are case-sensitive (age, Age and AGE are three different variables)

#### 5. There is no limit on the length of the variable name

#### 6. A variable name cannot contain spaces

#### 7. The variable name cannot be any Go keywords
