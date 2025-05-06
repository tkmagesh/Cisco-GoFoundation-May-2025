package main

import "fmt"

func main() {
	// product => id, Name, Cost, Units
	/*
		var p struct {
			Id    int
			Name  string
			Cost  float64
			Units int
		}
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 10
		p.Units = 20
	*/

	var p struct {
		Id    int
		Name  string
		Cost  float64
		Units int
	} = struct {
		Id    int
		Name  string
		Cost  float64
		Units int
	}{
		Id:    100,
		Name:  "Pen",
		Cost:  10,
		Units: 20,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	// fmt.Printf("%+v\n", p)
	fmt.Println(formatProduct(p))
}

func formatProduct(p struct {
	Id    int
	Name  string
	Cost  float64
	Units int
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d", p.Id, p.Name, p.Cost, p.Units)
}
