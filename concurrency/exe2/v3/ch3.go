package main

import (
	"fmt"
	"sync"
)

const GAP = 5

// introduce WaitGroup and for-range loop over channels
func main() {
	fmt.Println("start ch3")

	var wg sync.WaitGroup

	total := make(chan int, 2)
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, ((i-1)*GAP)+1, (i * GAP), total, &wg)
	}

	wg.Wait()

	close(total)
	for sum := range total {
		fmt.Println("sum:", sum)
	}

	fmt.Println("end")
}

func worker(wid int, start int, end int, total chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\twid:%v, start:%v, end:%v\n", wid, start, end)

	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}

	total <- sum
}
