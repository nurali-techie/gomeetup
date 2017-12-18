package nofunc

import "fmt"

type Counter interface {
	Next() int
}

type IntCounter struct {
	counter int
}

func (c *IntCounter) Next() int {
	c.counter = c.counter + 1
	return c.counter
}

func NoFuncDemo() {
	fmt.Println("NoFuncDemo - counter using interface way")

	counter := IntCounter{counter:10}

	fmt.Println(counter.Next())
	fmt.Println(counter.Next())
}
