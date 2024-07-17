/*
interfaces are implicitly implemented
*/

package main

import (
	"fmt"
	"math"
)

// ver 1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ver 2.0
type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("x has no Area() implementation")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("x has no Area() implementation")
	}
}
*/

/*
func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}
*/

type AreaFinder interface{ Area() float64 }

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// ver 4.0
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

type PerimeterFinder interface{ Perimeter() float64 }

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// ver 5.0

/*
	func PrintStats(x interface {
		AreaFinder
		PerimeterFinder
	}) {

		PrintArea(x)      //x must implement AreaFinder interface
		PrintPerimeter(x) //x must implement PerimeterFinder interface
	}
*/
type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintStats(x ShapeStatsFinder) {
	PrintArea(x)      //x must implement AreaFinder interface
	PrintPerimeter(x) //x must implement PerimeterFinder interface
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{Length: 10, Breadth: 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)

}
