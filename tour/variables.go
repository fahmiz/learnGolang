package main

import "fmt"

var c, python, java bool
var a, b int = 1, 2

func main() {
	var i int
	fmt.Println("variabel yang belum diset nilai:")
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println("variabel yang sudah diset nilai:")
	fmt.Println(a, b, c, python, java)

	k := 3
	fmt.Println("deklarasi variabel singkat (hanya bisa di dalam fungsi saja):")
	fmt.Println(k)
}
