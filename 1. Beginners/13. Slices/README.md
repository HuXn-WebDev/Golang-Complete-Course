# Slices

### Slices are also used to store multiple values of the same type in a single variable, however unlike arrays, the length of a slice can grow and shrink as you see fit.

## There are several ways to create a slice 👇

### 1. Using the []datatype{values} format

### 2. Create a slice from an array

### 3. Using the make() function

```go
// name := []datatype{values}
// name := []int{}

package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}
```

# Make() Method

### The make function will create a zeroed array and return a slice referencing an array. This is a great way to create a dynamically sized array. To create a slice using the make function, we need to specify three arguments: type, length and the capacity.

```go
package main
import "fmt"
func main() {
    slice := make([]string, 3, 5)
    fmt.Println("Length", len(slice))
    fmt.Println("Capacity", cap(slice))
    fmt.Println(slice)
}
```
