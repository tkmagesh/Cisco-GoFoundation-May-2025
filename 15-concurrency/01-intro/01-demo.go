package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of f1() through the "go scheduler" to be executed in FUTURE
	f2()

	// block the execution so that the scheduler looks for other goroutines scheduled and executte them (cooperative multitasking)
	time.Sleep(1 * time.Second)

	// the application is shutdown when the main() function completes
	fmt.Println("Done")
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
