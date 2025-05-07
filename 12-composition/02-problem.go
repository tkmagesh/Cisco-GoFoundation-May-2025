/*
Create the respective types to represent a Shopping Cart

Product
	Id, Name, UnitCost
LineItem
	Product, Units
ShoppingCart
	Multiple LineItems

1. User should be able to add items to the cart ()
2. Calculate the overall cart value and print
*/

/*
TODO:
1. Create a module
2. Create a "models" package that has all the models (Product, LineItem, ShoppingCart)
3. Use the models in the "main.main()"
4. Add "Remove Line Item" functionality
5. Add "Update Line Item Quantity" functionality
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

func (p Product) Format() string {
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

func (li LineItem) Format() string {
	return fmt.Sprintf("%s, Units = %d, Item Cost = %0.2f", li.Prod.Format(), li.Units, li.ItemCost())
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

func (sc ShoppingCart) Format() string {
	var sb strings.Builder
	for _, lineItem := range sc.LineItems {
		sb.WriteString(fmt.Sprintf("%s\n", lineItem.Format()))
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
	fmt.Println(cart.Format())
}
