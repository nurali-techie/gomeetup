package file

import (
	"fmt"
	"io"
	"os"
)

func WriteFileDemo(fname string) {
	fmt.Println("WriteFileDemo ..")

	// create file
	f, _ := os.Create(fname)
	defer f.Close()

	// write to file
	f.Write([]byte("abcd efgh ijkl mnop "))
}

func ReadFileDemo(fname string) {
	fmt.Println("ReadFileDemo ..")

	// open file
	f, _ := os.Open(fname)
	defer f.Close()

	// read from file
	readFromFile(f)
}

func readFromFile(r io.Reader) {
	b := make([]byte, 20)
	r.Read(b)
	fmt.Println(string(b))
}
