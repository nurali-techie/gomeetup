package str

import "fmt"

func StringReaderDemo() {
	fmt.Println("StringReaderDemo ..")

	// init StringReader with LONG string
	sr := &StringReader{"ABCD EFGH IJKL MNOP "}

	// read byte[10]
	b := make([]byte, 10)

	sr.Read(b)
	fmt.Println(string(b))

	sr.Read(b)
	fmt.Println(string(b))
}

// type StringReader, Src
type StringReader struct {
	Src string
}

// Read(), copy
func (sr *StringReader) Read(p []byte) (n int, err error) {
	n = len(p)
	copy(p, sr.Src[:n])
	sr.Src = sr.Src[n:]
	return n, nil
}
