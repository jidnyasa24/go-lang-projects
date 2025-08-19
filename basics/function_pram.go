// Question: Show short form for parameters of same type.
package main

import "fmt"

// Instead of (x int, y int), we write (x, y int)
func multiply(x, y int) int {
    return x * y
}

func main() {
    fmt.Println("Product of 3 and 4:", multiply(3, 4))
}
