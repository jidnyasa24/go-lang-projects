package main

import "fmt"

// Describe takes an empty interface and prints its type and value
func Describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func RunStory7() {
	fmt.Println("Story 7 - Empty Interface Example")

	Describe(42)           // integer
	Describe("hello Go")   // string
	Describe(3.14)         // float
	Describe([]int{1, 2})  // slice
}
