package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// r := strings.NewReader("Hello, Readerrrr!")
	// r := strings.NewReader("AaBbCcDd, EeFfGgHhIiJj!")
	r := strings.NewReader("abcdefghABCDEFGH!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err=%v b(ASCII) = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
