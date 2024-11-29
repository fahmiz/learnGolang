package main

import "fmt"

type Vertex struct {
	X int
	Y float32
	Z string
}

func main() {
	v := Vertex{1, 2.5, "ah"}
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v.Z)
}
