package main

import "fmt"

// Q: How can we make a function that works with multiple types (like int, string, float)?
// A: By using type parameters with a constraint such as `comparable`.

// Index finds the index of x in slice s.
// T must be a type that supports == and != (because of `comparable`).
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1 // return -1 if not found
}

func main() {
	// Q: Can the same function work for both integers and strings?
	// A: Yes! That’s the power of generics.

	ints := []int{10, 20, 30, 40}
	fmt.Println("Index of 30:", Index(ints, 30))

	strs := []string{"apple", "banana", "cherry"}
	fmt.Println("Index of 'banana':", Index(strs, "banana"))
}
