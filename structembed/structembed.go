package structembed

import "fmt"

// Go supports embedding structs and interfaces to express a more seamless composition of types

type Status struct {
	joined bool
}

type Student struct {
	ID    int
	Name  string
	Grade string
	Status
}

func StructEmbed() {
	fmt.Println("Embedding this struct!!")
}
