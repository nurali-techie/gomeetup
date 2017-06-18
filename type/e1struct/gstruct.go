package main

import "fmt"

/*
Typical OO language

class Person {
}
class Emp extends Person {
	Dept dept
}
class Dept {
}

*/

type Person struct {
	name string
	city string
}

type Emp struct {
	Person	// only composition (embedding)
	salary int
	Dept	// Person and Dept both embedded (kine of multiple inheritance)
}

type Dept struct {
	code string
}

func (p *Person) relocate(dest string) {
	fmt.Printf("%v relocated from %v to %v\n", p.name, p.city, dest)
	p.city = dest
}

func (p Person) move(dest string) {
	fmt.Printf("%v relocated from %v to %v\n", p.name, p.city, dest)
	p.city = dest
}

func main() {
	d1()
	d2()
}

func d2() {
	fmt.Println("\ndemo2")
	e1 := Emp{Person{"ali", "a1"}, 1000, Dept{"dev"}}
	e1.relocate("a2")
	fmt.Println(e1.code)
}

func d1() {
	fmt.Println("demo1")

	p1 := Person{"ali", "a1"}
	p2 := &Person{"jay", "j1"}

	p1.move("a2")
	p2.move("j2")	// syntex suger - (*p2).move()
	fmt.Println("City NOT changed - ", p1, p2)
	fmt.Println()

	p1.relocate("a3")	// syntax suger - (&p1).relocate()
	p2.relocate("j3")
	fmt.Println("City changed - ", p1, p2)
}
