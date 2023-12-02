# Arrays

### Arrays is a data structure which is use to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

### Arrays are 0 index based.

```go
package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```
