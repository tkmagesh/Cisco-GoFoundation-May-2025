package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}

	wg.Add(1) //increment the wait group counter by 1
	go f1(wg) //scheduling the execution of f1() through the "go scheduler" to be executed in FUTURE

	f2()

	// block the execution of this function until the counter becomes 0 (default = 0)
	wg.Wait()

	// the application is shutdown when the main() function completes
	fmt.Println("Done")
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() //decrement the waitgroup counter by 1
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
