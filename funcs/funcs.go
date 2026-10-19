package funcs

import "fmt"

func GoFuncs() {
	fmt.Println("Go Functions!!")
	// functions are central in Go.
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Here is the num sum:", sum(1, 4))
	fmt.Println("Here is the array sum:", sumArr(nums))

	// call the vals function here
	a, b := vals()
	fmt.Println(a, b)

	// ignore unused values
	_, c := vals()
	fmt.Println(c)

	// call the variadic function getSum here
	getSum(nums...)

	// call the intSeq closure function
	nextInt := intSeq()
	fmt.Println(nextInt())

	// call the fact recursive function
	fmt.Println(fact(5))
}

// sum of two numbers
func sum(x, y int) int {
	return x + y
}

// array sum
func sumArr(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// Multiple Return Values: This feature of Go is mainly used in idiomatic Go, for example to return both result and error values from a function
func vals() (int, int) {
	return 3, 7
}

// variadic functions
// They can be called with any number of trailing arguments. For example fmt.Println is a common variadic function.
func getSum(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Closures
// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a function inline without having to name it.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.

// Recursion
// Go supports recursive functions.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
