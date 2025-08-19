
// Q1: What is a goroutine?
// A goroutine is like a lightweight thread managed by Go runtime.
// You can start one with the `go` keyword.
// Example: `go f(x, y, z)` will run `f` in a new goroutine.


package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// This runs in the main goroutine.
	say("main function")

	// This runs in a separate goroutine.
	go say("hello from goroutine")

	// Keep main alive for a bit so goroutine can finish.
	time.Sleep(500 * time.Millisecond)
}
