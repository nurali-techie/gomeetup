package main

import (
	"fmt"
	"sync"
)

var total int = 0
var mutex sync.Mutex

var wg sync.WaitGroup

// this version has not only parallelism but also concurrency
// worker() has leap year logic, which has been splitted into two steps
// now even though you have only one worker(), both worker() and split() will run concurrently

// Rob pike talk - Concurrency is not Parallelism
// https://talks.golang.org/2012/waza.slide#1
// https://talks.golang.org/2012/waza.slide#12	- problem
// https://talks.golang.org/2012/waza.slide#27  - solution: concurrently with parallelism

// why this solution is concurrent - because leap year logic splitted and now running simultaneously
// why this solution is parallel - because more than one worker() running simultaneously
func main() {
	fmt.Println("start leap3")

	wg.Add(3)
	go worker(1, 1000, 1500)
	go worker(2, 1501, 2000)
	go worker(3, 2001, 2017)
	wg.Wait()

	fmt.Println("end, total:", total)
}

func worker(wid int, start int, end int) {
	defer wg.Done()
	fmt.Printf("\twid:%v, start:%v, end:%v\n", wid, start, end)

	years := make(chan int, 10)
	done := make(chan bool)
	go split(years, done)

	for i := start; i <= end; i++ {
		if i%4 == 0 {
			no := i
			years <- no
		}
	}

	close(years)
	<-done
}

func split(years chan int, done chan bool) {
	count := 0
	for year := range years {
		leapYear := false
		if year%400 == 0 {
			leapYear = true
		} else if year%100 == 0 {
			leapYear = false
		} else {
			leapYear = true
		}

		if leapYear {
			count++
		}
	}

	mutex.Lock()
	total += count
	mutex.Unlock()

	done <- true
}
