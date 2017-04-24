package main

import "fmt"

func main() {
	fmt.Println("start ch1")

	total := make(chan int)
	go worker(1, 100, total)

	fmt.Println("end, total:", <-total)
}

func worker(start int, end int, total chan int) {
	fmt.Println("\tstart:", start, ", end:", end)

	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}

	total <- sum
}
