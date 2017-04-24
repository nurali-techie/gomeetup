package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start ch2")

	total := make(chan int, 2)
	go worker(1, 1, 5, total)
	go worker(2, 6, 10, total)

	time.Sleep(time.Second * 1)

	fmt.Println("total1:", <-total)
	fmt.Println("total2:", <-total)

	fmt.Println("end")
}

func worker(wid int, start int, end int, total chan int) {
	fmt.Printf("\twid:%v, start:%v, end:%v\n", wid, start, end)

	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}

	fmt.Println("\tsum:", sum)
	total <- sum
	fmt.Println("\tend, wid:", wid)

}
