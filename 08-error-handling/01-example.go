package main

import (
	"errors"
	"fmt"
)

func main() {
	divisor := 7
	// divisor := 0
	q, r, err := divide(100, divisor)
	if err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d, remainder = %d\n", divisor, q, r)
	} else {
		fmt.Println("Error occurred :", err)
	}
}

/*
func divide(multiplier, divisor int) (int, int, error) {
	if divisor == 0 {
		err := errors.New("divisor cannot be 0")
		return 0, 0, err
	}
	quotient, remainder := multiplier/divisor, multiplier%divisor
	return quotient, remainder, nil
}
*/

// Named results
func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	quotient, remainder = multiplier/divisor, multiplier%divisor
	return
}
