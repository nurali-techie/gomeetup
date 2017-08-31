package slicedemo

import (
	"fmt"
)

func Test21() {
	fmt.Println("In slicedemo.Test21")
	// slice it and change one element, range
	s1 := []int{10, 20, 30, 40, 50}
	pslice("s1", s1)

	// first arg 	-> from which index new slice will start
	// second arg 	-> lenth of new slice = second_arg - first_arg
	s2 := s1[1:3]
	pslice("s2", s2) // check what part of s1 is referred by s2

	s2[0] = 21
	pslice("s1", s1) // check s1 got 21 anywhere
	pslice("s2", s2)
	fmt.Println()
}

func Test22() {
	fmt.Println("In slicedemo.Test22")
	// play with two slice
	s1 := []int{10, 20, 30, 40, 50}
	pslice("t1.s1", s1)

	s2 := s1[1:3]
	pslice("t2.s2", s2)

	s2 = append(s2, 60)
	pslice("t3.s1", s1) // check s1 got 60 anywhere
	pslice("t3.s2", s2)

	s2 = append(s2, 70, 80)
	pslice("t4.s1", s1) // check what elements of s1 changed
	pslice("t4.s2", s2) // check capacity of s2
	fmt.Println()

	// ** Output explained **

	// Output:
	// In slicedemo.Test22
	// t1.s1=[10 20 30 40 50], len=5, cap=5
	// t2.s2=[20 30], len=2, cap=4
	// t3.s1=[10 20 30 60 50], len=5, cap=5
	// t3.s2=[20 30 60], len=3, cap=4
	// t4.s1=[10 20 30 60 50], len=5, cap=5
	// t4.s2=[20 30 60 70 80], len=5, cap=8

	// at t2 time, s2 will get len=2, cap=4.
	// at t3 time, s1 got the changes
	// at t4 time, s1 not get any changes
	// at t4 time, for s2 new array will be allocated with double capacity (i.e. cap=8)
	// conclusion - when appending new element to s2, if it has enough capacity then it will append to exising slice
	// if s2 does not has enough capacity then new array will be allocated in memory for s2
}

func Test23() {
	fmt.Println("In slicedemo.Test23")
	// play with three slice
	s1 := []int{10, 20, 30, 40, 50}
	pslice("t1.s1", s1)

	s2 := s1[1:3]
	pslice("t2.s2", s2)

	s3 := append(s2, 60, 70, 80)
	pslice("t3.s3", s3)

	s3[0] = 21
	pslice("t4.s1", s1) // check any changes to s1
	pslice("t4.s2", s2) // check any changes to s2
	pslice("t4.s3", s3) // check capacity of s3
	fmt.Println()
}

func Test24() {
	fmt.Println("In slicedemo.Test24")
	// slice with three args
	s1 := []int{10, 20, 30, 40, 50}
	pslice("t1.s1", s1)

	// first arg 	-> from which index new slice will start
	// second arg 	-> lenth of new slice = second_arg - first_arg
	// third arg 	-> capacity of new slice = third_arg - first_arg
	// use third arg to make sure future append to s2 will result in new array allocation for s2 to make sure that s2 detach from s1
	s2 := s1[1:3:3]
	pslice("t2.s2", s2)

	// comment 's2 = append(s2, 60)' and check output for 's2[0] = 21'
	s2 = append(s2, 60) // as capacity and length was same, any new append will allocate new array for s2
	pslice("t3.s1", s1)
	pslice("t3.s2", s2) // check capacity of s2

	s2[0] = 21
	pslice("t4.s1", s1) // check any changes to s1
	pslice("t4.s2", s2)
	fmt.Println()
}
