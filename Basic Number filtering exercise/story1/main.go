package main

import (
	"fmt"
)

func main() {
	// Sample input
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Call function to get even numbers
	evens := getEvenNumbers(numbers)

	// Print output
	fmt.Println("Even numbers:", evens)
}

// Function to filter even numbers
func getEvenNumbers(nums []int) []int {
	var evens []int
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	return evens
}
