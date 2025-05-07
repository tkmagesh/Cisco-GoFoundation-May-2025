package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	primesCh := genPrimes(2, 100)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

// producer
func genPrimes(start, end int) chan int {
	primesCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				primesCh <- no
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(primesCh)
	}()
	return primesCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
