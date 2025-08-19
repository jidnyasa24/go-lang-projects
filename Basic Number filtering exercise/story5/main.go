// Story 5: Given a list of integers, write a program to return only the even and multiples of 5 from this list.
package main

import (
	"fmt"
)

func main() {
	// Sample input
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Call function to get evens and multiples of 5
	result := getEvenAndMultiplesOfFive(numbers)

	// Print output
	fmt.Println("Even numbers & multiples of 5:", result)
}

// Function to filter even numbers and multiples of 5
func getEvenAndMultiplesOfFive(nums []int) []int {
	var result []int
	for _, n := range nums {
		if n%2 == 0 && n%5 == 0 {
			result = append(result, n)
		}
	}
	return result
}
