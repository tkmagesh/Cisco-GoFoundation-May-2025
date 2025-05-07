package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	x = 100
	x = true
	x = 99.87
	x = "Labore pariatur ullamco eiusmod consequat esse exercitation et cupidatat exercitation consequat et Lorem ad laboris."
	fmt.Println(x)

	x = 200
	// y := x * 2 //=> compilation error

	x = "Sint officia culpa velit excepteur."

	// y := x.(int) * 2 //=> results in a panic (unsafe)
	// fmt.Println(y)

	fmt.Println("type assertion using if-else")
	// safe approach (type assertion)
	// x = "Sint officia culpa velit excepteur."
	x = 200
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x does not have a value of type int")
	}
	fmt.Println("\n================================================")
	fmt.Println("type assertion using type-switch")
	// type assertion using "type switch"
	// x = "Sint officia culpa velit excepteur."
	// x = 200
	// x = false
	// x = 99.87
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, x * 0.01 =", val*0.01)
	default:
		fmt.Println("x is of unknown type")
	}

}
