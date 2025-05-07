/*
Modify the below in such a way that the "increment()" is executed "concurrently"
Ensure that the "count" is printed AFTER all the increment function executions are completed
*/
package main

import "fmt"

var count int

func main() {
	for range 300 {
		increment()
	}
	fmt.Println("count :", count)
}

func increment() {
	count += 1
}
