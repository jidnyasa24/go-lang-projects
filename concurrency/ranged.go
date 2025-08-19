// Q4: How do we use range and close with channels?
package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // close channel when done sending
}

func main() {
	c := make(chan int, 10)
	go fibonacci(10, c)

	// range automatically ends when channel is closed
	for i := range c {
		fmt.Println(i)
	}
}
