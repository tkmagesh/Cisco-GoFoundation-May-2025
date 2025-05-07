/*
Struct Methods
*/

package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Cost  float64
	Units int
}

// Methods (= functions with receivers)
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d", p.Id, p.Name, p.Cost, p.Units)
}

/* Receiver has to be a pointer if the struct state has to be changed */
func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func main() {

	p := Product{
		Id:    100,
		Name:  "Pen",
		Cost:  10,
		Units: 20,
	}

	fmt.Println("Before applying 10% discount")

	// fmt.Println(formatProduct(p))
	fmt.Println(p.Format())

	// applyDiscount(&p, 10)
	p.ApplyDiscount(10)

	fmt.Println("Before applying 10% discount")
	// fmt.Println(formatProduct(p))
	fmt.Println(p.Format())

}
