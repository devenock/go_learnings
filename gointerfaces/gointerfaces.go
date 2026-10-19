package gointerfaces

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perimeter() int
}

type Rectangle struct {
	Width  int
	Height int
}

type Circle struct {
	Radius float64
}

// methods
func (r Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

func (r Rectangle) Perimeter() int {
	return 2*r.Width + 2*r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// if a variable has an interface type, then we can call methods that are in the names interface.
func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}

func GoInterfaces() {
	fmt.Println("Go Interfaces!!")
	// Interfaces are named collections of method signatures
	// To implement an interface in Go, we just need to implement all the methods in the interface
	rect := Rectangle{Width: 2, Height: 5}
	// circle := Circle{Radius: 7}

	measure(rect)
	// measure(circle)
}
