package arraydemo

import "fmt"

func Test21() {
	fmt.Println("In arraydemo.Test21")
	// array copy
	a1 := [5]int{10, 20, 30, 40, 50}
	parr("t1.a1", a1)
	a2 := a1 // this will make deep copy
	parr("t2.a1", a1)
	parr("t2.a2", a2)
	a2[2] = 31 // this will NOT change a1[2]
	parr("t3.a1", a1)
	parr("t3.a2", a2)

	// diff/same len copy
	a3 := [5]int{10, 20, 30, 40, 50}
	a4 := [5]int{} // a4 len must be same as a3, try "a4 := [4]int{}" and check error
	a4 = a3
	parr("t4.a3", a3)
	parr("t4.a4", a4)
	fmt.Println()
}

func Test22() {
	fmt.Println("In arraydemo.Test22")
	// pass array to func change1()
	a1 := [3]string{"red", "green", "blue"}
	fmt.Printf("t1.a1=%v\n", a1)
	change1(a1)
	fmt.Printf("t2.a1=%v\n", a1) // check value of a1[1]
	fmt.Println()
}

func Test23() {
	fmt.Println("In arraydemo.Test23")
	// pass array to func change2()
	a1 := [3]string{"red", "green", "blue"}
	fmt.Printf("t1.a1=%v\n", a1)
	change2(&a1)                 // observe &
	fmt.Printf("t2.a1=%v\n", a1) // check value of a1[1]
	fmt.Println()
}

func change1(colors [3]string) {
	colors[1] = "GREEN"
}

func change2(colors *[3]string) {
	colors[1] = "GREEN"
}
