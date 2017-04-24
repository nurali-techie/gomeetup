package main

import (
	"fmt"
	"sync"
)

var total int = 0
var wg sync.WaitGroup

// run this multiple times and you will get diff result in each run.
// this version shows there is a problem in this logic which creates race condition.
func main() {
	fmt.Println("start cnt1")

	wg.Add(2)

	go counter(100000)
	go counter(100000)

	wg.Wait()

	fmt.Println("end, total:", total)
}

func counter(times int) {
	defer wg.Done()
	for i := 1; i <= times; i++ {
		total += 1
	}
}
