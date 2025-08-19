


// Q6: Methods and pointer indirection
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)   // Go auto converts to (&v).Scale(2)
	p := &v
	p.Scale(10)  // works directly
	fmt.Println(v)
}
