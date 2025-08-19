// Q7: How do we check if two binary trees store the same values using concurrency?
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t and sends all values to channel ch
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines if t1 and t2 store the same values
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Tree(1) vs Tree(1):", Same(tree.New(1), tree.New(1)))
	fmt.Println("Tree(1) vs Tree(2):", Same(tree.New(1), tree.New(2)))
}
