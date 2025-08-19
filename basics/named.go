// Question: Demonstrate named return values with "naked" return.
package main

import "fmt"

// split splits a number into two parts
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // naked return uses named variables
}

func main() {
    a, b := split(17)
    fmt.Println("Split of 17:", a, b)
}
