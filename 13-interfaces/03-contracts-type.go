package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

// utilities
/*
func PrintArea(x interface{}) {
	// Version 1.0
	// switch val := x.(type) {
	// case Circle:
	// 	fmt.Println("Area :", val.Area())
	// case Rectangle:
	// 	fmt.Println("Area :", val.Area())
	// default:
	// 	fmt.Println("Not a type with Area method")
	// }

	// Version 2.0
	switch val := x.(type) {
	case interface {
		Area() float32
	}:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Not a type with Area method")
	}

}
*/
type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShapeStats(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShapeStats(r)
}
