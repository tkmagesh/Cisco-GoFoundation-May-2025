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

/* Composition */

type PerishableProduct struct {
	Product
	Expiry string
}

// Overriding the methods of the base (Product)
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {
	milk := PerishableProduct{
		Expiry: "2 Days",
		Product: Product{
			Id:    200,
			Name:  "Milk",
			Cost:  40,
			Units: 2,
		},
	}

	// Accessing the members
	fmt.Println(milk.Expiry)
	// fmt.Println(milk.Product.Id, milk.Product.Name)
	fmt.Println(milk.Id, milk.Name)

	// Accessing the methods of the base (Product) type
	fmt.Println(milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println(milk.Format())
}
