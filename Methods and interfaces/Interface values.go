// Q11: Interface values
// ----------------------
// Under the hood, interface values are like (value, type).
// When you call a method on an interface, it runs the method
// of the underlying concrete type.

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println("Method M() called with:", t.S)
}

func main() {
	var i I
	i = T{"Hello Jidnyasa!"}
	i.M()
}
