/*
Modify the below in such a way that the "increment()" is executed "concurrently"
Ensure that the "count" is printed AFTER all the increment function executions are completed
*/
package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		count += 1 //critical path
	}
	mutex.Unlock()
}
