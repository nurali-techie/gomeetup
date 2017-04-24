package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("start route1")

	fmt.Println("before:", runtime.NumGoroutine()) // print no_of_goroutines

	go worker() // will run in diff goroutine
	fmt.Println("after1:", runtime.NumGoroutine())

	go fmt.Println("hello world from main") // call func from other package in diff goroutine
	fmt.Println("after2:", runtime.NumGoroutine())

	time.Sleep(time.Second * 1)

	fmt.Println("end")
}

func worker() {
	fmt.Println("\tstarting work")
	fmt.Println("\twork done")
}
