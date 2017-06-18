package main

import "fmt"

// built-in types (int) can be re-defined as normal type, after that methods [stage()] can be attached
type Age int

type Person struct {
	name string
	Age
}

func (p Person) stage() string {
	if p.Age >= 18 {
		return "adult"
	}
	return "teen"
}

func main() {
	p1 := Person{"ali", Age(30)}
	fmt.Println(p1.stage())

	p2 := Person{"jay", Age(12)}
	fmt.Println(p2.stage())
}
