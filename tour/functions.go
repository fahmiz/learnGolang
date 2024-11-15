package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func pangkat3(x uint8) uint8 {
	return x * x * x
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(pangkat3(3))
}
