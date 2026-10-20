package goerrors

import (
	"errors"
	"fmt"
)

func GoErrors() {
	fmt.Println("Go Errors!!")
	// In Go, its idiomatic to communicate errors in an explicit, separate return value.
	myNum := 52
	fmt.Println(f(myNum))

	showMarks()
}

// a function that returns an error
func f(arg int) (int, error) {
	if arg == 42 {
		// The New function creates errors whose only content is a text message.
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil //a nil value indicates that there was no error
}

var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 42 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

// You can also use an inline error check(it's idiomatic)
var marks = []int{3, 5, 7, 8, 9}

func showMarks() {
	for _, val := range marks {
		if r, e := f(val); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}
}
