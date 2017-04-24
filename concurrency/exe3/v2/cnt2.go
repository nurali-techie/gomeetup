package main

import (
	"fmt"
	"sync"
)

var total int = 0
var mutex sync.Mutex

var wg sync.WaitGroup

// solution using mutex
// Go proverb - Do not communicate by sharing memory; instead, share memory by communicating.
// https://golang.org/doc/effective_go.html#concurrency
// ** This shows - communicate by sharing memory (but do not do this way) **
func main() {
	fmt.Println("start cnt2")

	wg.Add(2)

	go counter(10000)
	go counter(10000)

	wg.Wait()

	fmt.Println("end, total:", total)
}

func counter(times int) {
	defer wg.Done()
	for i := 1; i <= times; i++ {
		// solve race condition just using Mutex lock/unlock
		mutex.Lock()
		total += 1
		mutex.Unlock()
	}
}
