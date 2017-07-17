package count

import (
	"fmt"
	"io"
)

func WordCountDemo(r io.Reader) {
	fmt.Println("WordCountDemo ..")

	// init wcnt = 0
	var wcnt int = 0

	// read using byte[10]
	b := make([]byte, 10)

	r.Read(b)
	wcnt = wcnt + countSpace(b)

	r.Read(b)
	wcnt = wcnt + countSpace(b)

	fmt.Println("wcnt=", wcnt)
}

func countSpace(b []byte) int {
	var spaceCnt int = 0
	for _, ch := range b {
		if ch == 32 {
			spaceCnt++
		}
	}
	return spaceCnt
}
