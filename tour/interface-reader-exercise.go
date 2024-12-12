package main

import (
	"fmt"
	"io"
)

type MyReader struct{}

// Implementasikan metode Read untuk MyReader
func (r MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A' // Isi setiap byte di buffer dengan karakter 'A'
	}
	return len(p), nil // Kembalikan panjang buffer yang diisi
}

func main() {
	r := MyReader{}

	// Gunakan io.LimitReader untuk membatasi jumlah byte yang dibaca
	lr := io.LimitReader(r, 20)

	// Gunakan io.ReadAll untuk membaca semua data dari LimitReader
	data, _ := io.ReadAll(lr)

	fmt.Printf("Output: %q\n", data) // Output: "AAAAAAAAAAAAAAAA"
}
