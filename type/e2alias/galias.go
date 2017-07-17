package main

import (
	"fmt"
)

// built-in types (int) can be re-defined as normal type, after that methods [stage()] can be attached
type Age int

type Person struct {
	name string
	Age
}

func (a Age) stage() string {
	if a >= 18 {
		return "adult"
	}
	return "teen"
}

func main() {
	p1 := Person{"ali", Age(30)}
	fmt.Printf("%v is %v\n", p1.name, p1.stage())

	p2 := Person{"jay", Age(12)}
	fmt.Printf("%v is %v\n", p2.name, p2.stage())
}
