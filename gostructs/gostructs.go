package gostructs

import "fmt"

type person struct {
	name string
	age  int
}

func GoStructs() {
	fmt.Println("Go structs!!")
	// Go structs are typed collection of fields. It is how go represents structured data without needing classes or inheritance.
	// The struct fields are laid out contigously in memory in the order you declare them.
	// A struct total size is the sum of its field sizes plus any padding, which is why grouping fields of similar size together can sometimes shrink a struct's memory footprint.
	// Every struct field has a well defined zero-value which means a struct is always ready to use the moment it's declared and there is no "uninitialized" state to guard against.
	// A struct have value semantics, meaning assigning a strut to a new variable, or passing it to a function, copies every field.
	// Useful for grouping data together to form records.

	// initiate an instance of person
	p := person{name: "Mike", age: 6}
	fmt.Println("New person:", p)

	// access a struct field using dot(.)
	fmt.Println("Mike's age:", p.age)

	// call newPerson
	fmt.Println("My newest person:", newPerson("Brian"))
}

// use a function to initialize a new person
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 65
	return &p
}
