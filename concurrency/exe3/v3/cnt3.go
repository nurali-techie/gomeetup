package main

import (
	"fmt"
	"sync"
)

//var total int = 0
var wg sync.WaitGroup

// solve using channels. this is the way go is promoting to do synchronization.
// Go proverb - Do not communicate by sharing memory; instead, share memory by communicating.
// https://golang.org/doc/effective_go.html#concurrency
// ** This shows - share memory by communicating **
func main() {
	fmt.Println("start cnt3")

	wg.Add(2)

	total := make(chan int, 1)
	total <- 0

	go counter(1, 1000000, total)
	go counter(2, 1000000, total)

	wg.Wait()

	fmt.Println("end, total:", <-total)
}

func counter(wid int, times int, total chan int) {
	defer wg.Done()
	fmt.Println("\twid:", wid)
	for i := 1; i <= times; i++ {
		//total += 1
		// incr value in channel
		temp := <-total
		temp = temp + 1
		total <- temp
	}
}
