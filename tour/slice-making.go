package main

import "fmt"

func main() {
	x := make([]int, 5, 5)
	fmt.Println("x := make([]int, 5, 5)")
	printSlice("x", x)

	a := make([]int, 5)
	fmt.Println("a := make([]int, 5)")
	printSlice("a", a)

	b := make([]int, 0, 5)
	fmt.Println("b := make([]int, 0, 5)")
	printSlice("b", b)

	c := b[:2]
	fmt.Println("c := b[:2]")
	printSlice("c", c)

	d := c[2:5]
	fmt.Println("d := c[2:5]")
	printSlice("d", d)

	e := c[0:5]
	fmt.Println("e := c[0:5]")
	printSlice("e", e)

	f := e[2:5]
	fmt.Println("f := e[2:5]")
	printSlice("f", f)

	g := d[2:3]
	fmt.Println("g := d[2:3]")
	printSlice("g", g)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
	fmt.Println()
}
