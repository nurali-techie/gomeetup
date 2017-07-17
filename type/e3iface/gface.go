package main

import "fmt"

type Animal interface{}	// empty interface, accepts all

type Walker interface {
	Walk()
}

type Flyer interface {
	Fly()
}

type Dog struct {
}

type Eagle struct {
}

func (d Dog) Walk() {
	fmt.Println("dog can walk")
}

func (e Eagle) Walk() {
	fmt.Println("eagle can walk")
}

func (e Eagle) Fly() {
	fmt.Println("eagle can fly")
}

func main() {
	d1()
	d2()
	d3()
	d4()
}

func d4() {
	fmt.Println("\nDemo4 - type assertion with switch")
	skill(Dog{})
	skill(Eagle{})
}

func skill(a Animal) {
	switch animal := a.(type) {
	case Flyer:
		animal.Fly()
	case Walker:
		animal.Walk()
	}
}

func d3() {
	fmt.Println("\nDemo3 - type assertion / type casting")
	flyMe(Eagle{})
	flyMe(Dog{})
}

func flyMe(a Animal) {
	//f := a.(Flyer)	// this will type-cast but it's unsafe
	if f, ok := a.(Flyer); ok {	// safe type casting
		f.Fly()
	} else {
		fmt.Println("can't fly")
	}
}

func d2() {
	fmt.Println("\nDemo2 - Polimorphism")
	walkMe(Dog{})
	walkMe(Eagle{})
}

func walkMe(w Walker) {
	w.Walk()
}

func d1() {
	fmt.Println("Demo1 - use interface on LHS")
	var w Walker = Dog{}
	w.Walk()
}
