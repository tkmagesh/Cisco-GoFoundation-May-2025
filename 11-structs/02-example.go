package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Cost  float64
	Units int
	// Category string
}

func main() {

	/*
		var p Product = Product{
			Id:    100,
			Name:  "Pen",
			Cost:  10,
			Units: 20,
		}
	*/

	/*
		var p = Product{
			Id:    100,
			Name:  "Pen",
			Cost:  10,
			Units: 20,
		}
	*/

	p := Product{
		Id:    100,
		Name:  "Pen",
		Cost:  10,
		Units: 20,
	}

	// The following is not advisable (requires all the attribute values to be mentioned)
	// p := Product{100, "Pen", 10, 20}

	fmt.Println("Before applying 10% discount")
	fmt.Println(formatProduct(p))
	applyDiscount(&p, 10)
	fmt.Println("Before applying 10% discount")
	fmt.Println(formatProduct(p))

	// struct instances are "values"
	/*
		x := Product{200, "Laptop", 10000, 5}
		y := x // a copy of x is created (coz structs are values)
		y.Cost = y.Cost - 1000
		fmt.Println(x)
		fmt.Println(y)
	*/

	/*
		x := Product{200, "Laptop", 10000, 5}
		y := Product{200, "Laptop", 10000, 5}
		fmt.Println("x == y ? :", x == y)
	*/

	// pointers
	/*
		x := Product{200, "Laptop", 10000, 5}
		var xPtr *Product
		xPtr = &x
		fmt.Println(xPtr.Id)
	*/

}

func formatProduct(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d", p.Id, p.Name, p.Cost, p.Units)
}

func applyDiscount(p *Product, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
