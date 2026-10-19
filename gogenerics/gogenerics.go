package gogenerics //type parameters

import "fmt"

func GoGenerics() {
	fmt.Println("Welcome to Go Generics!!")
	// Generics, introduced in Go 1.18, let you write functions and types that work with
	// multiple specific types while keeping type safety.
	// It allows functions and types to operate on multiple specific types without the need for code duplication.

	// KEY FEATURES
	// 1. Type Parameters: These are placeholder for types, defined in square brackets after a function or type name.
	// 2. Type cONSTRAINT: They limit which types can be used with a generic function or type, ensuring that only compatible types are allowed
	// 3. Type Inference: Go can often infer the type parameters from the function arguments, simplifying the syntax for developers.

	// BENEFITS OF USING GENERICS
	// 1. Code Reuse: Write one function or type that works for many types, reducing redundancy.
	// 2. Type Safety: cATCH TYPE ERRORS AT COMPILE TIME, PREVENTING RUNTIME ISSUES.
	// 3. Cleaner Code: Avoid repetitive code and risky type casting, leading to more maintainable code.

	// calling the non-generic functions
	int64Map := map[string]int64{
		"Math":    23,
		"English": 4,
		"Biology": 5,
	}
	myInt64Sum := SumInts(int64Map)
	fmt.Println(myInt64Sum)

	float64Map := map[string]float64{
		"Average": 45.2,
		"Median":  20.5,
		"Mode":    50.2,
	}

	myFloat64Sum := SumFloats(float64Map)
	fmt.Println(myFloat64Sum)

	// call the generic function
	fmt.Println(SumIntsOrFloats[string, int64](int64Map))
	fmt.Println(SumIntsOrFloats[string, float64](float64Map))

	// You can as well ignore the type arguments since the compiler can always infer them
	fmt.Println(SumIntsOrFloats(float64Map))
}

// non-generic functions
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// With Generics, you can just write one function
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, n := range m {
		s += n
	}
	return s
}

// EXAMPLE OF A GENERIC FUNCTION WHEN EXPLICITLY DECLARING A type constraint

// declare the type constraint as an interface
type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, n := range m {
		s += n
	}
	return s
}

// we can also have a struct with a generic type
type element[T any] struct {
	next *element[T]
	val  T
}
type List[T any] struct {
	head, tail *element[T]
}
