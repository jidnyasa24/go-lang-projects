// Question: Demonstrate exported names in Go.
// In math package, Pi (capital P) is exported, pi (small p) is not.
package main

import (
    "fmt"
    "math"
)

func main() {
    // Using math.Pi (exported, works fine)
    fmt.Println("Value of Pi is:", math.Pi)
}
