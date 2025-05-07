package main

import (
	"fmt"
	"math"
)

/* Circle */
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

/* Rectangle */
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

/* Ver 2.0 */
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

/* Ver 3.0 */
// Feel free to add any methods to existing shapes (Circle, Rectangle, Square)

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

type PerimeterFinder interface {
	Perimeter() float64
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/*
func PrintShape(x interface {
	AreaFinder
	PerimeterFinder
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x ShapeStatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 16}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 14, Width: 15}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)

	s := Square{Side: 16}
	/*
		PrintArea(s)
		PrintPerimeter(s)
	*/
	PrintShape(s)

	// PrintArea(100)
}
