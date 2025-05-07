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

/*
func PrintArea(x interface{}) {
	// fmt.Println("Area :", x.Area())
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("invalid argument (argument does not have Area() method)")
	}
}
*/

/*
func PrintArea(x interface{}) {
	// fmt.Println("Area :", x.Area())
	switch val := x.(type) {
	case interface{ Area() float64 }: // case "any object with Area() method"
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("invalid argument (argument does not have Area() method)")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
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
func PrintPerimeter(x /* ? */) {
	/* ? */
}

func main() {
	c := Circle{Radius: 16}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 14, Width: 15}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
	PrintPerimeter(r)

	s := Square{Side: 16}
	PrintArea(s)
	PrintPerimeter(s)

	// PrintArea(100)
}
