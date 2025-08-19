// Question: Simple function with two int parameters.
package main

import "fmt"

// Function takes two ints and returns their sum
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println("Sum of 5 and 6:", add(5, 6))
}
