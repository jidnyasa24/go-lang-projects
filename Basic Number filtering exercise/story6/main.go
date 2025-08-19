// Story 6: Given a list of integers, write a program to return only the odd and multiples of 3 greater than 10 from this list.
package main

import (
	"fmt"
)

func main() {
	// Sample input
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Call function to get odd & multiples of 3 greater than 10
	result := getOddAndMultiplesOfThreeGreaterThanTen(numbers)

	// Print output
	fmt.Println("Odd & multiples of 3 > 10:", result)
}

// Function to filter odd numbers & multiples of 3 greater than 10
func getOddAndMultiplesOfThreeGreaterThanTen(nums []int) []int {
	var result []int
	for _, n := range nums {
		if n > 10 && n%2 != 0 && n%3 == 0 {
			result = append(result, n)
		}
	}
	return result
}
