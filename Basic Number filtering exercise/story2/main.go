package main

import (
	"fmt"
)

func main() {
	// Sample input
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Get odd numbers
	odds := getOddNumbers(numbers)

	// Print result
	fmt.Println("Odd numbers:", odds)
}

// Function to filter odd numbers
func getOddNumbers(nums []int) []int {
	var odds []int
	for _, n := range nums {
		if n%2 != 0 {
			odds = append(odds, n)
		}
	}
	return odds
}
