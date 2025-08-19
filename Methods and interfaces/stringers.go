// Q16: Exercise - Stringers
// --------------------------
// Task: Make the IPAddr type implement fmt.Stringer so it prints like "1.2.3.4".
// Example: IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

package main

import (
	"fmt"
	"strings"
)

// Define IPAddr as a slice of 4 bytes
type IPAddr [4]byte

// Implement Stringer interface
func (ip IPAddr) String() string {
	// Convert each byte to string and join with "."
	parts := make([]string, len(ip))
	for i, b := range ip {
		parts[i] = fmt.Sprint(b)
	}
	return strings.Join(parts, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
