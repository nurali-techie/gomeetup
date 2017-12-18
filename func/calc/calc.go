package calc

import "fmt"

type Operation func(int, int) int

func CalcDemo() {
	calcDemo1()
	calcDemo2()
	calcDemo3()
	calcDemo4()
}

func calcDemo1() {
	fmt.Println("CalcDemo1 - direct calling function with name")
	a := 20
	b := 10
	fmt.Println("add=", add(a,b))
	fmt.Println("sub=", sub(a,b))
}

func calcDemo2() {
	fmt.Println("CalcDemo2 - calling function with variable")
	a := 20
	b := 10
	fmt.Println("add=", calc(a,b, add))
	fmt.Println("sub=", calc(a,b, sub))
}

func calcDemo3() {
	fmt.Println("CalcDemo3 - calling function with variable from map")
	oper := make(map[string]Operation)
	oper["+"] = add
	oper["-"] = sub

	a := 20
	b := 10
	fmt.Println("add=", calc(a,b, oper["+"]))
	fmt.Println("sub=", calc(a,b, oper["-"]))
}

func calcDemo4() {
	fmt.Println("CalcDemo4 - calling function with return value from other function")
	a := 20
	b := 10
	fmt.Println("add=", calc(a,b, findOper("+")))
	fmt.Println("sub=", calc(a,b, findOper("-")))
}

func calc(x, y int, op Operation) int {
	return op(x,y)
}

func findOper(symbol string) Operation {
	switch symbol {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func add(x, y int) int {
	return x+y
}

func sub(x, y int) int {
	return x-y
}
