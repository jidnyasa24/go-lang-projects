// Q15: Type switches
// -------------------
// Like a switch, but for types inside an interface.
// Helps to check what actual type is stored.

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v*2)
	case string:
		fmt.Println("String length:", len(v))
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}

func main() {
	do(21)
	do("Jidnyasa")
	do(true)
}
