package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p) // Baca data dari pembaca asli
	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') || (p[i] >= 'a' && p[i] <= 'z') { //'A' = 65 ascii
			if p[i] >= 'A' && p[i] <= 'Z' { // == if p[i] >= 65 && p[i] <= 90 {
				// fmt.Println("p before[", i, "]:", p[i])
				p[i] = 'A' + (p[i]-'A'+13)%26
				// fmt.Println("p after[", i, "]:", p[i])
			} else if p[i] >= 'a' && p[i] <= 'z' {
				// fmt.Println("p[", i, "]:", p[i])
				p[i] = 'a' + (p[i]-'a'+13)%26
				// fmt.Println("p[", i, "]:", p[i])
			}
		}
	}
	return n, err
}

func main() {
	fmt.Println("Exercise: rot13Reader")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()
	s = strings.NewReader("AZaznopqabcd0123")
	r = rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()
	fmt.Println((1 + 3) % 3) // % 3 menjaga siklus 1-3, angka 4 kembali ke 1
}

// https://go.dev/tour/methods/23
