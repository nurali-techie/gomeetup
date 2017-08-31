package slicedemo

import (
	"fmt"
)

func Test11() {
	fmt.Println("In slicedemo.Test11")
	// init direct
	a1 := []int{10, 20, 30, 40, 50}
	pslice("a1", a1)

	// init make, len
	a2 := make([]int, 5)
	pslice("a2", a2)

	// init make, len, cap
	a3 := make([]int, 3, 5)
	pslice("a3", a3)

	// init direct with last index
	a4 := []int{10, 20, 4: 50}
	pslice("a4", a4)
	fmt.Println()
}

func pslice(name string, slice []int) {
	fmt.Printf("%v=%v, len=%v, cap=%v\n", name, slice, len(slice), cap(slice))
}
