package main

import "fmt"

func add(x int, y int) int {
	fmt.Println("Menggunakan func add(x int, y int) int {}")
	return x + y
}

func add_too(x, y int) int {
	fmt.Println("Menggunakan func add_too(x, y int) int {}")
	fmt.Println("x dan y tipe data nya sama jadi bisa ditulis seperti ini saja")
	return x + y
}

func pangkat3(x uint8) uint8 {
	fmt.Println("Menggunakan func pangkat3(x uint8) uint8 {}")
	return x * x * x
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add_too(42, 13))
	fmt.Println(pangkat3(3))
}
