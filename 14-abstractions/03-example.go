/*
Using standard library interfaces (fmt.Stringer interface)
*/

package main

import (
	"fmt"
	"strings"
)

/* Product Type */
type Product struct {
	Id       int
	Name     string
	UnitCost float64
}

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, UnitCost = %0.2f", p.Id, p.Name, p.UnitCost)
}

/* LineItem Type */
type LineItem struct {
	Prod  Product
	Units int
}

func (li LineItem) ItemCost() float64 {
	return float64(li.Units) * li.Prod.UnitCost
}

// fmt.Stringer interface implementation
func (li LineItem) String() string {
	return fmt.Sprintf("%s, Units = %d, Item Cost = %0.2f", li.Prod, li.Units, li.ItemCost())
}

/* ShoppingCart Type */
type ShoppingCart struct {
	LineItems []LineItem
}

func (sc *ShoppingCart) AddItem(prod Product, units int) {
	lineItem := LineItem{Prod: prod, Units: units}
	sc.LineItems = append(sc.LineItems, lineItem)
}

func (sc ShoppingCart) GetCartValue() float64 {
	cartValue := float64(0)
	for _, lineItem := range sc.LineItems {
		cartValue += lineItem.ItemCost()
	}
	return cartValue
}

// fmt.Stringer interface implementation
func (sc ShoppingCart) String() string {
	var sb strings.Builder
	for _, lineItem := range sc.LineItems {
		sb.WriteString(fmt.Sprintf("%s\n", lineItem))
	}
	sb.WriteString("============================================================================================\n")
	sb.WriteString(fmt.Sprintf("Cart Value : %0.2f\n", sc.GetCartValue()))
	return sb.String()
}

func main() {
	pen := Product{Id: 100, Name: "Pen", UnitCost: 10}
	pencil := Product{Id: 101, Name: "Pencil", UnitCost: 5}
	marker := Product{Id: 102, Name: "Marker", UnitCost: 50}

	cart := ShoppingCart{}
	cart.AddItem(pen, 10)
	cart.AddItem(pencil, 30)
	cart.AddItem(marker, 4)
	fmt.Println(cart)
}
