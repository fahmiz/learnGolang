package main

import "fmt"

func main() {
	fmt.Println("Menambahkan elemen ke slice dengan append")
	fmt.Println("var s []int")
	var s []int
	printSlice(s)

	fmt.Println("append bisa bekerja pada slice yang nil.")
	fmt.Println("s = append(s, 0)")
	s = append(s, 0)
	printSlice(s)

	fmt.Println("Slice bertambah seperlunya.")
	fmt.Println("s = append(s, 1)")
	s = append(s, 1)
	s = append(s, 1)
	printSlice(s)

	fmt.Println("Kita juga bisa menambahkan lebih dari satu elemen sekaligus.")
	fmt.Println("s = append(s, 2, 3, 4)")
	s = append(s, 2, 3, 4)
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()
}
