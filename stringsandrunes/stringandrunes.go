package stringsandrunes

import (
	"fmt"
	"unicode/utf8"
)

func StringsAndRunes() {
	// a string in Go is a read-only slice of bytes
	// a character is a called a rune in Go. A rune is an integer that represents a Unicode code point.
	fmt.Println("strings and runes!!")
	const name = "Mike"

	// get the length of the string
	fmt.Println("Name length:", len(name))

	// loop through a string: The loop generates the hex values of all the bytes that constitute the code points in name
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	// to count how many runes/characters we have in a string, we use the utf package
	fmt.Println("Rune count:", utf8.RuneCountInString(name))
}
