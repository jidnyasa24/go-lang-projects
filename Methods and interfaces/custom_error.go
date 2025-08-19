// Q18: Exercise - Custom Errors with Sqrt
// ----------------------------------------
// Modify Sqrt function to return an error when input is negative.
// Create custom error type ErrNegativeSqrt that implements error interface.

package main

import (
	"fmt"
	"math"
)

// Custom error type
type ErrNegativeSqrt float64

// Implement the Error() method for ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt function that returns (value, error)
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	// Valid input
	result, err := Sqrt(16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sqrt(16) =", result)
	}

	// Invalid input
	result, err = Sqrt(-2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sqrt(-2) =", result)
	}
}
