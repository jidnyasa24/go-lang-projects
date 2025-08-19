// Q17: Errors in Go
// ------------------
// Go uses error values to indicate failure states.
// A nil error means success, a non-nil error means failure.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Try converting a string to integer
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("Couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

	// Try with invalid input
	_, err2 := strconv.Atoi("not-a-number")
	if err2 != nil {
		fmt.Printf("Conversion failed: %v\n", err2)
	}
}
