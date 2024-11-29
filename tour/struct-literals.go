package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // memiliki tipe Vertex
	v2 = Vertex{X: 1}  // Y:0 adalah implisit
	v3 = Vertex{}      // X:0 dan Y:0
	v4 = Vertex{Y: 1}  // X:0 adalah implisit
	p  = &Vertex{3, 4} // memiliki tipe *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3, v4)
}
