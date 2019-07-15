package main

import (
	"fmt"
	"os"

	"github.com/nurali-techie/gomeetup/reader/buf"
	"github.com/nurali-techie/gomeetup/reader/count"
	"github.com/nurali-techie/gomeetup/reader/file"
	"github.com/nurali-techie/gomeetup/reader/str"
)

func main() {
	fmt.Println("start test")

	demo1()

	demo2()

	demo3()

	demo4()

	demo5()

	demo6()

	fmt.Println("end")
}

func demo1() {
	fmt.Println("demo1 - write to file and read from same file using io.Reader")
	file.WriteFileDemo("data.txt")
	file.ReadFileDemo("data.txt")
	fmt.Println()
}

func demo2() {
	fmt.Println("demo2 - count no of Words in given file using io.Reader")
	f, _ := os.Open("data.txt")
	defer f.Close()
	count.WordCountDemo(f)
	fmt.Println()
}

func demo3() {
	fmt.Println("demo3 - read file using BufReader")
	f, _ := os.Open("data.txt")
	defer f.Close()
	buf.BufReaderDemo(f)
	fmt.Println()
}

func demo4() {
	fmt.Println("demo4 - read string using StringReader")
	str.StringReaderDemo()
	fmt.Println()
}

func demo5() {
	fmt.Println("demo5 - count no of words in given string using StringReader")
	sr := &str.StringReader{"AAAA BBBB CCCC DDDD "}
	count.WordCountDemo(sr)
	fmt.Println()
}

func demo6() {
	fmt.Println("demo6 - count no of words in given BufReader which uses StringReader")
	sr := &str.StringReader{"AAAA BBBB CCCC DDDD "}
	br := &buf.BufReader{sr, make([]byte, 0)}
	count.WordCountDemo(br)
	fmt.Println()
}
