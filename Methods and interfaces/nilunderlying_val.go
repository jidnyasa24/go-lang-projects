// Q12: Interface values with nil underlying values
// -------------------------------------------------
// If an interface contains a nil value, the method can still be called,
// but it must handle nil gracefully.

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("Called M() on a nil receiver")
		return
	}
	fmt.Println("Method M() with value:", t.S)
}

func main() {
	var i I
	var t *T
	i = t
	i.M() // works gracefully even though t is nil
}
