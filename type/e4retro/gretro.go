package main

import "fmt"

// --- Third party code. Modification NOT possible.
type PaidZip struct {
}

func (z PaidZip) Zip() {
	fmt.Println("Paid to zip")
}

// ---

// --- My code
type MyZip struct {
}

func (z MyZip) Zip() {
	fmt.Println("Free to zip")
}

// ---

func main() {
	z1 := PaidZip{}
	send(z1)

	z2 := MyZip{}
	send(z2)
}

// Just add new interface and use in send() to take PaidZip as well as MyZip without modification to PaidZip code
type Zipper interface {
	Zip()
}

func send(z Zipper) {
	// lot of logic
	z.Zip()
	// send
}

//func send2(z MyZip) {
//	// lot of logic
//	z.zip()
//	// send
//}
