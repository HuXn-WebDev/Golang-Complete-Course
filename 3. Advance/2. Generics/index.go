// Without Generics
// package main

// import "fmt"

// // Function printNumber
// func printNumber(item, defaultValue int) (int, int) {
//     return item, defaultValue
// }

// // Function printString
// func printString(item, defaultValue string) (string, string) {
//     return item, defaultValue
// }

// // Function printBoolean
// func printBoolean(item, defaultValue bool) (bool, bool) {
//     return item, defaultValue
// }

// func main() {
//     // Example usage
//     num, _ := printNumber(42, 0)
//     fmt.Println(num) // Outputs: (42, 0)

//     str, _ := printString("hello", "world")
//     fmt.Println(str) // Outputs: (hello, world)

//     boolean, _ := printBoolean(true, false)
//     fmt.Println(boolean) // Outputs: (true, false)
// }
// -----------------------------------

// With Generics
package main

import "fmt"

// generic function printItem
func printItem[T any](item, defaultValue T) (T, T) {
    return item, defaultValue
}

func main() {
    // Example usage
    num, _ := printItem(42, 0)
    fmt.Println(num) // Outputs: (42, 0)

    str, _ := printItem("hello", "world")
    fmt.Println(str) // Outputs: (hello, world)

    boolean, _ := printItem(true, false)
    fmt.Println(boolean) // Outputs: (true, false)
}

// -----------------------------------
// --- Example 2 ---

// package main

// import "fmt"

// // A generic function to compare two values of any type
// func Compare[T comparable](a, b T) bool {
// 	return a == b
// }

// func main() {
// 	// Compare integers
// 	resultInt := Compare(10, 20)
// 	fmt.Println("Are the integers equal?", resultInt)

// 	// Compare strings
// 	resultStr := Compare("hello", "world")
// 	fmt.Println("Are the strings equal?", resultStr)
// }
