// Question: Function that returns multiple values.
package main

import "fmt"

// swap returns two strings in reverse order
func swap(a, b string) (string, string) {
    return b, a
}

func main() {
    first, second := swap("hello", "world")
    fmt.Println("After swap:", first, second)
}
