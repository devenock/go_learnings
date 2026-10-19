package gorange

import "fmt"

func GoRange() {
	fmt.Println("Go Range!!")
	// range iterates over elements in a variety of built-in data structures
	nums := []int{2, 4, 5, 7, 9, 1}
	sum := 0
	// loop using range
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Range sum:", sum)

	// range over slices and arrays provides values and index. But you can ignore like above using the underscore sign(_)
	marks := [...]int{24, 20, 12, 26, 15}
	for i, m := range marks {
		fmt.Printf("Index: %v and value: %v\n", i, m)
	}

	// range on map iterates over key/value pairs
	myMap := make(map[string]int)
	myMap["students"] = 14
	myMap["average"] = 20
	myMap["parents"] = 30
	myMap["buses"] = 2
	fmt.Println("Dorm map:", myMap)
	// use range to loop
	for key, val := range myMap {
		fmt.Println("key:", key, "val:", val)
	}

	// you can also just iterate over the keys alone
	for key := range myMap {
		fmt.Println("key here:", key)
	}

	// range over strings iterate over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itseld.
	for _, char := range "school" {
		fmt.Println("char:", char)
	}
}
