package constants

import (
	"fmt"
)

func GoConstants() {
	// Go supports constants of character, string, boolean and numeric values.
	// You declare a constant variables value using the const keyword
	const name string = "Mellon"
	// Numeric constants has got no type until it's given one, such as by an explicit conversion.
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}
