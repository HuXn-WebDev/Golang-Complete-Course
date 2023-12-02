# Structs (Structures)

### A struct is used to create a collection of members of different data types, into a single variable.

```go
package main
import ("fmt")

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
  var userOne Person

    userOne.name = "HuXn"
    userOne.age = 18
    userOne.job = "Programmer"
    userOne.salary = 40000
    fmt.Println(userOne)
    fmt.Println("My name is", userOne.name, "I'm", userOne.age, "Years old", "My Profession is", userOne.job, "My salary is", userOne.salary)
}
```
