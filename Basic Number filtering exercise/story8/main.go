package main

import (
	"fmt"
)

func main() {
	// Sample input
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Example 1: prime, greater than 15, multiple of 5
	res1 := filterAny(numbers, prime, greaterThan(15), multipleOf(5))
	fmt.Println("Output 1:", res1) // Expect [2 3 5 7 10 11 13 15 16 17 18 19 20]

	// Example 2: less than 6, multiple of 3
	res2 := filterAny(numbers, lessThan(6), multipleOf(3))
	fmt.Println("Output 2:", res2) // Expect [1 2 3 4 5 6 9 12 15 18]
}
