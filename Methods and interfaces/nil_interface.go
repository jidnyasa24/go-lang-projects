// Q13: Nil interface values
// --------------------------
// A nil interface holds no value and no type.
// Calling a method on it will cause a runtime error.

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	if i == nil {
		fmt.Println("i is nil and has no type")
	}
	// i.M() // Uncommenting this will panic!
}
