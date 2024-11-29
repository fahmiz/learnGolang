package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	q := v //q bukan pointer
	q.X = 1e3
	fmt.Println(v) // v tidak berubah

	p := &v //p pointer
	p.X = 1e3
	fmt.Println(v) // v berubah
}
