package main

import "fmt"

func main() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	defer fmt.Println()
	defer fmt.Println("world")

	fmt.Println("Perintah defer menunda eksekusi dari sebuah fungsi sampai fungsi yang melingkupinya selesai.")
	fmt.Println()
	fmt.Println("Argumen untuk pemanggilan defer dievaluasi langsung, tapi pemanggilan fungsi tidak dieksekusi sampai fungsi yang melingkupinya selesai.")
	fmt.Println()
	fmt.Println("hello")
}
