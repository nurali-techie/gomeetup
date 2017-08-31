package slicedemo

import "fmt"

func Test31() {
	fmt.Println("In slicedemo.Test31")
	// copy(slice) .. compare copy array
	s1 := []int{10, 20, 30, 40, 50}
	s2 := make([]int, 2) // change len to 2 and check output
	pslice("t1.s1", s1)
	pslice("t1.s2", s2)

	copy(s2, s1)
	pslice("t2.s1", s1)
	pslice("t2.s2", s2)

	s2[0] = 11
	pslice("t3.s1", s1) // check s1[0] changed
	pslice("t3.s2", s2)
	fmt.Println()
}

func Test32() {
	fmt.Println("In slicedemo.Test32")
	// change1 changes
	s1 := []int{10, 20, 30, 40, 50}
	pslice("t1.s1", s1)

	change1(s1)
	pslice("t2.s1", s1)

	// change2 no change
	s2 := []int{10, 20, 30, 40, 50}
	pslice("t3.s2", s2)

	change2(s2)
	pslice("t4.s2", s2)
	fmt.Println()

	// conclusion - slice is always pass by reference
	// changes done in change1(s1) will be reflected to s1 with one exception as shown in change2()
	// in case of change2(s2)
	// append(s, 60) create new array so it will detach from slice s2
	// and further changes in slice s will not be reflected on s2
}

func change1(s []int) {
	// modify all elements
	for i := 0; i < len(s); i++ {
		s[i] = s[i] + 1
	}
	pslice("change1.s", s)
}

func change2(s []int) {
	// append one new element
	s = append(s, 60)
	for i := 0; i < len(s); i++ {
		s[i] = s[i] + 1
	}
	pslice("change2.s", s)
}
