// Q2: How do channels allow goroutines to communicate safely?
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// Send the result into the channel
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from channel

	fmt.Println("Results:", x, y, "Total:", x+y)
}
