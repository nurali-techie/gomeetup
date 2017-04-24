package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// this version uses goroutines
// we have introduce parallelism (but not concurrency) in this solution
// here two worker() running in parallel
func main() {
	fmt.Println("start leap2")

	wg.Add(2)
	go worker(1, 1000, 1500)
	go worker(2, 1501, 2000)
	wg.Wait()

	fmt.Println("end")
}

func worker(wid int, start int, end int) {
	defer wg.Done()
	fmt.Printf("\twid:%v, start:%v, end:%v\n", wid, start, end)

	count := 0
	for i := start; i <= end; i++ {
		leapYear := false
		if i%4 == 0 {
			leapYear = split(i)
		}

		if leapYear {
			count++
		}
	}
	fmt.Printf("\twid:%v, count:%v\n", wid, count)
}

func split(year int) bool {
	leapYear := false
	if year%400 == 0 {
		leapYear = true
	} else if year%100 == 0 {
		leapYear = false
	} else {
		leapYear = true
	}
	return leapYear
}
