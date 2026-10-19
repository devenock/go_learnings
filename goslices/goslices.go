package goslices

import (
	"fmt"
	"slices"
)

func GoSlices() {
	fmt.Println("Slices in Go!")
	marks := []int{2, 4, 5, 6, 7, 9}
	// add item to end of the slice
	marks = append(marks, 8)

	// creating a slice using make
	ages := make([]int, 7)
	fmt.Println("my ages here:", ages)

	// getting length and capacity of a slice(always same by default)
	fmt.Println("length is:", len(ages), "and capacity is:", cap(ages))

	// slice operator
	l := marks[:3]
	fmt.Println("my sliced marks here:", l)

	// slice comparisons
	if slices.Equal(ages, marks) {
		fmt.Println("ages == marks")
	}
	fmt.Println(marks)
}
