package main

import "fmt"

func main() {
	// Declaring interface
	// Implementing interface
	// Empty interface
	// Multiple interfaces
	// Type assertion
	// Type switch
	// Embedding interfaces
	// Pointer vs Value receiver
	// Interface comparison

	var s Shape
	fmt.Printf("value of s %v\n", s)
	fmt.Printf("type of s %T\n", s)

	// Interface have 2 type
	// static and dynamic

	//Type assertion is not only used to check if an interface has a concrete value of some given type but also to convert a given variable of an interface type to a different interface type

	rect := Rect{
		width:  132,
		height: 133,
	}
	s = rect
	fmt.Printf("value of s %v\n", s)
	fmt.Printf("type of s %T\n", s)

}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}
