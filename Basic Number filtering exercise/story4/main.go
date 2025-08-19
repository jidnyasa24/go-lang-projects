// Story 4: Given a list of integers, write a program to return only the odd prime numbers from this list.
package main

import (
	"fmt"
)

func main() {
	// Sample input
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Call function to get odd primes
	oddPrimes := getOddPrimeNumbers(numbers)

	// Print output
	fmt.Println("Odd prime numbers:", oddPrimes)
}

// Helper function to check if a number is prime
func isPrime(n int) bool {
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

// Function to filter odd primes
func getOddPrimeNumbers(nums []int) []int {
	var oddPrimes []int
	for _, n := range nums {
		if isPrime(n) && n%2 != 0 {
			oddPrimes = append(oddPrimes, n)
		}
	}
	return oddPrimes
}
