package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(p []byte) (int, error) {
	// Isi buffer p dengan karakter 'A'
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil // Kembalikan jumlah byte yang diisi dan nil untuk error
}

func main() {
	reader.Validate(MyReader{})
}

// run in https://go.dev/tour/methods/22
