package main

import (
	"net/http"
	"fmt"
	"io"
	"bytes"
)

/*
if you want something like this ..
type I interface {
	M(in) out
}

.. you can do it -> interface way
func API(I)

.. or you can do it -> higher-order function way
func API(func M(in) out)

.. again you want interface with only one method
.. at that time you can decide whether you want to go
1. interface-way
2. higher-order function way
3. both way
*/

func funcHandler(res http.ResponseWriter, req *http.Request) {
	// handler logic
}

//from http pakcage from standard library
//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
type MyHandler struct {
}
func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// handler logic
}

func main() {
	httpHandlerBothWay()
	ioReaderInterfaceWay()
}

func httpHandlerBothWay() {
	// you can provide handler function to http in both way
	http.HandleFunc("/", funcHandler)	// higher-order function way
	http.Handle("/", MyHandler{})	// interface way
}

func ioReaderInterfaceWay() {
	var reader io.Reader = &bytes.Buffer{}
	fmt.Println(reader)
	var rc io.ReadCloser = nil
	fmt.Println(rc)
}
