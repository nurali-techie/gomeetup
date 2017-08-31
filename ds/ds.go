package main

import (
	"github.com/gomeetup/ds/arraydemo"
	"github.com/gomeetup/ds/slicedemo"
)

func main() {
	arrayTest()
	sliceTest()
}

func arrayTest() {
	arraydemo.Test11()
	arraydemo.Test12()

	arraydemo.Test21()
	arraydemo.Test22()
	arraydemo.Test23()
}

func sliceTest() {
	slicedemo.Test11()

	slicedemo.Test21()
	slicedemo.Test22()
	slicedemo.Test23()
	slicedemo.Test24()

	slicedemo.Test31()
	slicedemo.Test32()
}
