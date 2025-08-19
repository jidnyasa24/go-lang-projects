// Q3: What happens when we use buffered channels?
package main

import "fmt"

func main() {
	// Create a buffered channel with capacity 2
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 3 // uncommenting this will cause deadlock because buffer is full

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
