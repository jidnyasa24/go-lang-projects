/*

Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
*/


// Question: Demonstrate using packages in Go.
// This program imports "fmt" (for printing) and "math/rand" (for random numbers).

package main

// Importing packages
import (
    "fmt"      // for printing to console
    "math/rand" // for random number generation
)

func main() {
    // Using fmt.Println to print a random number between 0 and 99
    fmt.Println("A random number:", rand.Intn(100))
}
