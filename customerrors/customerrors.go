package customerrors

import "fmt"

// It is possible to define custom error types by implementing the Error() method on them.
// A custom error type always have the suffix "Error"

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string { // Adding this Error method makes argError implement the error interface.
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// return custom error
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func CustomErrors() {
	fmt.Println("Go Custom Errors")
}
