package main

import "fmt"

func main() {
	fmt.Println("Nilai kosong dari slice adalah nil.")
	fmt.Println("Slice yang kosong memiliki panjang dan kapasitas 0, dan tidak memiliki array di dalamnya.")
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
