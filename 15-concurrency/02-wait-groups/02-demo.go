package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of f1() through the "go scheduler" to be executed in FUTURE
	f2()

	// Poorman's synchronization technique (DO NOT do this)
	time.Sleep(4 * time.Second)

	// the application is shutdown when the main() function completes
	fmt.Println("Done")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
