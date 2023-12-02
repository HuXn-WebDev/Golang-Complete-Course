# panic()

### Similar to exceptions raised at runtime when an error is encountered. panic() is either raised by the program itself when an unexpected error occurs or the programmer throws the exception on purpose for handling particular errors.

```go
package main

func employee(name *string, age int){
  if age > 65{
    panic("Age cannot be greater than retirement age")
  }

}

func main() {
  empName := "Samia"
  age := 75
  employee(&empName, age)
}
```
