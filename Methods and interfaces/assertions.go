// Q14: Type assertions

// Used to access the concrete value inside an interface.
// t := i.(T) will panic if i is not of type T.
// t, ok := i.(T) is the safe form (returns boolean).

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Direct assertion
	s := i.(string)
	fmt.Println("String value:", s)

	// Safe assertion
	safe, ok := i.(string)
	fmt.Println("Safe assertion:", safe, ok)

	// Wrong type
	num, ok := i.(float64)
	fmt.Println("Wrong type assertion gives:", num, ok)
}
