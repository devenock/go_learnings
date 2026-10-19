package gomethods

import "fmt"

// Go supports methods, which are functions with receivers. The receiver is what is defined before the function name.

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() int {
	return 2*r.Width + 2*r.Height
}

func GoMethods() {
	fmt.Println("Go Methods!!")
	rect := Rectangle{Width: 4, Height: 3}
	fmt.Println("Area here:", rect.Area())
	fmt.Println("Perimeter here:", rect.Perimeter())
}
