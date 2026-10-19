package values

import "fmt"

func GoValues() {
	// Go have various value types: strings, int, bool, floats

	// strings
	firstname := "Molly"
	lastname := "Munes"
	fmt.Println(firstname + " " + lastname)

	// int
	age := 20
	// float
	averageMarks := 78.5
	// bool
	isVerified := false
	fmt.Println(age, averageMarks, isVerified)
}
