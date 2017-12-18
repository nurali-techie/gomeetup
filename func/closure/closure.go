package closure

import "fmt"

func generator1() func() int {
	counter := 0
	return func() int {
		counter = counter + 1
		return counter
	}
}

func generator2(start int) func() int {
	counter := start
	return func() int {
		counter = counter + 1
		return counter
	}
}

func generator3(start, gap int) func() int {
	counter := start
	return func() int {
		counter = counter + gap
		return counter
	}
}

func generator4(start int) func(int) int {
	counter := start
	return func(gap int) int {
		counter = counter + gap
		return counter
	}
}

func ClosureDemo() {
	closureDemo1()
	closureDemo2()
	closureDemo3()
	closureDemo4()
}

func closureDemo1() {
	fmt.Println("ClosureDemo1 - no argument")
	increment := generator1()
	fmt.Println(increment())
	fmt.Println(increment())
}

func closureDemo2() {
	fmt.Println("ClosureDemo2 - argument for counter start value")
	increment := generator2(10)
	fmt.Println(increment())
	fmt.Println(increment())
}

func closureDemo3() {
	fmt.Println("ClosureDemo3 - argument for counter start with increment gap value")
	increment := generator3(10, 2)
	fmt.Println(increment())
	fmt.Println(increment())
}

func closureDemo4() {
	fmt.Println("ClosureDemo4 - argument for counter start with different gap value")
	increment := generator4(10)
	fmt.Println(increment(5))
	fmt.Println(increment(20))
}
