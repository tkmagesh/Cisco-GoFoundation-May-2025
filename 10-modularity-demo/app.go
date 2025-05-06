package main

import (
	"fmt"

	"github.com/tkmagesh/Cisco-GoFoundation-May-2025/10-modularity-demo/calculator"
	"github.com/tkmagesh/Cisco-GoFoundation-May-2025/10-modularity-demo/calculator/utils"
)

func run() {
	fmt.Println("Application executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Op Count :", calculator.GetOpCount())

	fmt.Println("Is 21 even ?:", utils.IsEven(21))
	fmt.Println("Is 21 odd ?:", utils.IsOdd(21))
}
