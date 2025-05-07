/*
Hint : Use strconv.Atoi() to convert integers in string format to integers
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10))               //=> 10
	fmt.Println(sum(10, 20))           //=> 30
	fmt.Println(sum(10, 20, "30", 40)) //=> 100
	fmt.Println(sum("abc", 10))        //=> 10
	fmt.Println(sum())                 //=> 0

	// TODO:
	fmt.Println(sum(10, 20, []int{30, 40}))                                    //=> 100
	fmt.Println(sum(10, 20, []interface{}{"30", 40}))                          //=> 100
	fmt.Println(sum(10, 20, []interface{}{"30", 40, []int{50, 60}}))           //=> 210
	fmt.Println(sum(10, 20, []interface{}{"30", 40, []interface{}{"50", 60}})) //=> 210
}

func sum(list ...interface{}) int {
	var result int
	for _, x := range list {
		switch val := x.(type) {
		case int:
			result += val
		case string:
			if no, err := strconv.Atoi(val); err == nil {
				result += no
			}
		}
	}
	return result
}
