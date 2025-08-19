// Story 7: Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that match ALL the conditions.
package main

import (
	"fmt"
)

// --- Modular condition functions ---
// Each of these takes an integer and returns a bool
func odd(n int) bool        { return n%2 != 0 }
func even(n int) bool       { return n%2 == 0 }
func greaterThan(k int) func(int) bool {
	return func(n int) bool { return n > k }
}
func lessThan(k int) func(int) bool {
	return func(n int) bool { return n < k }
}
func multipleOf(k int) func(int) bool {
	return func(n int) bool { return n%k == 0 }
}
func prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// --- Orchestration: ALL conditions must pass ---
func filterAll(nums []int, conditions ...func(int) bool) []int {
	var result []int
	for _, n := range nums {
		matches := true
		for _, cond := range conditions {
			if !cond(n) {
				matches = false
				break
			}
		}
		if matches {
			result = append(result, n)
		}
	}
	return result
}

func main() {
	// Sample input
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Example 1: odd, greater than 5, multiple of 3
	res1 := filterAll(numbers, odd, greaterThan(5), multipleOf(3))
	fmt.Println("Output 1:", res1) // Expect [9 15]

	// Example 2: even, less than 15, multiple of 3
	res2 := filterAll(numbers, even, lessThan(15), multipleOf(3))
	fmt.Println("Output 2:", res2) // Expect [6 12]
}
