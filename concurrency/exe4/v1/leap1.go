package main

import "fmt"

// find no_of_leap years in given range
// solved without concurrency
func main() {
	fmt.Println("start leap1")

	count := worker(1000, 2000)

	fmt.Println("end, count:", count)
}

func worker(start int, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		leapYear := false
		if i%4 == 0 {
			leapYear = filter(i)
		}

		if leapYear {
			count++
		}
	}
	return count
}

func filter(year int) bool {
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
