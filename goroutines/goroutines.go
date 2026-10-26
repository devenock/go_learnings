package goroutines

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution
// They are more efficient than traditional OS threads because they require less memoryand have faster startup time.
// Every Go program starts with a main goroutine and when it finishes execution, all other goroutines are terminated.
func f(from string) {
	for i := range 5 {
		fmt.Println(from, ":", i)
	}
}

func GoRoutines() {
	fmt.Println("Go Routines!!")

	// go routine invokation
	go f("Reading")

	// normal invokation
	f("school")

	// anonymous func goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
