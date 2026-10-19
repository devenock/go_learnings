package goiterators //added in v1.23

import (
	"fmt"
	"iter"
)

// the order of iteration over maps is randomized to prevent programmers from relying on any specific sequence
var marks = map[string]int{
	"Maths":     34,
	"English":   20,
	"Swahili":   45,
	"Chemistry": 80,
	"Business":  70,
}

// iterate over marks
func listMarks(m map[string]int) {
	for key, val := range m {
		fmt.Printf("Key: %s. Value: %d\n", key, val) // shows different order of results each time you run it. This design choice encourages developers to write code that does not depend on the order of map elements, leading to more robust and predictable programs.
	}
}

// range over string to get a rune(each character value)
func rangeOverString(n string) {
	for _, i := range n {
		fmt.Printf("CharCode: %v\n", i)
	}
}

// range over integers
func rangeOverInteger(n int) {
	for i := range n {
		fmt.Printf("Num: %d\n", i)
	}
}

// range over channels
func rangeOverChan() {
	ch := make(chan int, 4)
	ch <- 10
	ch <- 20
	close(ch)
	for i := range ch {
		fmt.Println("Chan", i)
	}
}

// iterators introduced in Go 1.23 makes it possible to create custom sequences that can be used directly with the range keyword.
// Here is a custom iterator that traverses a slice in reverse.
// The yield function is called for each element, allowing the range loop to iterate over the sequence seamlessly.
func Backward[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

// declare struct generics
type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

// a method that uses a generic
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Iterator in Go is a func with a special signature
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func GoIterators() {
	fmt.Println("Welcome to Go Iterators!!")
	// we can range over literally anything using iterators

	listMarks(marks)

	rangeOverString("Mike")

	rangeOverInteger(5)

	rangeOverChan()

	// range over iterators
	nums := []int{1, 2, 4, 5, 7, 8}
	for i, n := range Backward(nums) {
		fmt.Printf("Index %d: %d\n", i, n)
	}
}
