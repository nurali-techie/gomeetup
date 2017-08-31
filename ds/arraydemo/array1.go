package arraydemo

import (
	"fmt"
)

func Test11() {
	fmt.Println("In arraydemo.Test11")
	// default int
	a1 := [5]int{}
	parr("a1", a1)

	// direct init
	a2 := [5]int{10, 20, 30, 40, 50}
	parr("a2", a2)

	// assign value
	a2[2] = 31
	parr("a2", a2)
	fmt.Println()
}

func Test12() {
	fmt.Println("In arraydemo.Test12")
	// range
	a1 := [5]int{10, 20, 30, 40, 50}
	for ind, ele := range a1 {
		fmt.Printf("ind=%v, ele=%v\n", ind, ele)
	}
	fmt.Println()
}

func parr(name string, arr [5]int) {
	fmt.Printf("%v=%v, len=%v\n", name, arr, len(arr))
}
