/* functions + pointers */

package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("[main] Before incrementing, x = %d\n", x)
	fmt.Println("[main] Address of x :", &x)
	increment(&x)
	fmt.Printf("[main] After incrementing, x = %d\n", x)
}

func increment(no *int) {
	fmt.Println("[increment] Address of no :", no)
	*no += 1
}
