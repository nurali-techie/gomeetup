package buf

import (
	"fmt"
	"io"
)

// In this demo, we are reading 5 bytes in each Read() call but BufReader is reading 10 bytes in each Read() call
// So every two call we make, BufReader will make only one call
func BufReaderDemo(r io.Reader) {
	fmt.Println("BufReaderDemo ..")

	// init BufReader
	br := &BufReader{r, make([]byte, 0)}

	// read byte[5]
	b := make([]byte, 5)

	// BufReader will read only once for first time
	br.Read(b)
	fmt.Println(string(b))
	br.Read(b)
	fmt.Println(string(b))

	// BufReader will read only once for second time
	br.Read(b)
	fmt.Println(string(b))
	br.Read(b)
	fmt.Println(string(b))
}

// type BufReader, R, Buf
type BufReader struct {
	R   io.Reader
	Buf []byte
}

// Read() with byte[10], copy
func (br *BufReader) Read(p []byte) (n int, err error) {
	n = len(p)
	if len(br.Buf) == 0 {
		br.Buf = make([]byte, 10)
		fmt.Println("Reading by BufReader ..")
		br.R.Read(br.Buf)
	}

	copy(p, br.Buf[:n])
	br.Buf = br.Buf[n:]

	return n, nil
}
