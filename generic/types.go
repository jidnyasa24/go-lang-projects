package main

import "fmt"

// Q: What if we want a data structure (like a linked list) that works with any type?
// A: We can define a generic type with a type parameter.

// List is a singly linked list that can hold any type of value T.
type List[T any] struct {
	val  T
	next *List[T]
}

// PushFront inserts a new element at the beginning of the list.
func (l *List[T]) PushFront(v T) *List[T] {
	return &List[T]{val: v, next: l}
}

// Print displays all elements in the list.
func (l *List[T]) Print() {
	for n := l; n != nil; n = n.next {
		fmt.Print(n.val, " -> ")
	}
	fmt.Println("nil")
}

func main() {
	// Q: Can we create a list of integers? Yes.
	intList := &List[int]{val: 10}
	intList = intList.PushFront(20)
	intList = intList.PushFront(30)
	intList.Print() // 30 -> 20 -> 10 -> nil

	// Q: What about a list of strings? Works the same way!
	strList := &List[string]{val: "world"}
	strList = strList.PushFront("hello")
	strList.Print() // hello -> world -> nil
}
