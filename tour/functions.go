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

func swap(x, y string) (string, string) {
	fmt.Println("Menggunakan fungsi multiple result atau return > 1 nilai.")
	return y, x
}

func split(sum int) (x, y int) {
	fmt.Println("Menggunakan fungsi dengan naked return")
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add_too(42, 13))
	fmt.Println(pangkat3(3))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
