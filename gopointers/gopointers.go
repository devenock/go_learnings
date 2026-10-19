package gopointers

import "fmt"

func GoPointers() {
	fmt.Println("Go Pointers!!")
	// Go supports pointers, allowing you to pass references to values and records within your program

	// calling the functions
	i := 1
	fmt.Println("Initial value of i:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	// string pointers
	var name string = "Mike"
	var nikname *string = &name
	fmt.Println("name:", name, "nikname:", *nikname)
	// (*) operator is used to declare a pointer variable and to access the value stored in a pointer address(dereferencing operator)
	// (&) operator is used return the address of the variable or to access the address of a variable to a pointer(address-of operator)
}

// examples of a func that takes in a pointer argument and one that takes a non-pointer argument

// no pointer
func zeroval(ival int) {
	ival = 0
}

// pointer
func zeroptr(iptr *int) {
	*iptr = 0
}
