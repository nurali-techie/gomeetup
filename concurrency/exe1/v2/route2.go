package main

import (
	"fmt"
)

func main() {
	fmt.Println("start route2")

	done := make(chan bool)
	go worker(done)
	over := <-done // sync/communicate using channels

	fmt.Println("end, over:", over)
}

func worker(done chan bool) {
	fmt.Println("\tstarting work")
	fmt.Println("\twork done")
	done <- true
}
