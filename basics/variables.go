// Question: Show variable declaration with var keyword.
package main

import "fmt"

// Package level variable
var globalVar int = 100

func main() {
    var localVar int = 20
    fmt.Println("Global:", globalVar, "Local:", localVar)
}
