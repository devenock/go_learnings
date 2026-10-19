package gomaps

import (
	"fmt"
	"maps"
)

func GoMaps() {
	fmt.Println("Go Maps Starts Here!!")
	// Maps are Go's built-in associative data type(cometimes called hashes or dicts in other languages)
	// To create an empty map, use the built-in make:
	myMap := make(map[string]string)

	// you can set key/value pairs using typical name[key] = val syntax
	myMap["name"] = "Melon"
	myMap["title"] = "Engineer"
	myMap["department"] = "Technology"
	myMap["age"] = "12yrs"

	// get value using a key
	n := myMap["name"]
	fmt.Println("name:", n)

	// if you try getting a non-existing key, the zero value of that type is returned.
	m := myMap["marks"]
	fmt.Println("marks:", m)

	// get the map length using the len keyword
	mLen := len(myMap)
	fmt.Println("Map length:", mLen)

	// you can delete a map item using the delete method(pass in the map name as first argument and the key you want to remove)
	delete(myMap, "name")
	fmt.Println("After deletion", myMap)

	// you can also clear all the key/value from the map by using the clear method
	clear(myMap)
	fmt.Println("After clearing", myMap)

	// you can also declare and initialize a map at once
	studentMarks := map[string]int{"Mark": 91, "Yvone": 76, "Byron": 80}
	fmt.Println(studentMarks)

	// you can also check if two maps are equal using the Equal utility function
	marks := map[string]int{"Mark": 91, "Yvone": 76, "Byron": 80}
	if maps.Equal(studentMarks, marks) {
		fmt.Println("studentMarks == marks")
	}
	fmt.Println(myMap)
}
