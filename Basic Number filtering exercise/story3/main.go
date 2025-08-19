// Story 3: Given a list of integers, write a program to return only the prime numbers from this list.
package main

import (
	"fmt"
)

func main() {
	// Sample input
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Call function to get prime numbers
	primes := getPrimeNumbers(numbers)

	// Print output
	fmt.Println("Prime numbers:", primes)
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

// Function to filter prime numbers
func getPrimeNumbers(nums []int) []int {
	var primes []int
	for _, n := range nums {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	return primes
}
