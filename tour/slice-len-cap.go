package main

import "fmt"

func main() {
	fmt.Println("Panjang dan kapasitas dari sebuah slice s dapat diambil dengan menggunakan ekspresi len(s) dan cap(s).")
	fmt.Println("Panjang dari sebuah slice yaitu jumlah dari elemen yang dimilikinya.")
	fmt.Println("Kapasitas dari sebuah slice adalah jumlah dari elemen array yang dikandungnya, dihitung dari elemen pertama di dalam slice.")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	fmt.Println("Slice the slice to give it zero length. s = s[:0]")
	s = s[:0]
	printSlice(s)

	fmt.Println("Extend its length. s = s[:4]")
	s = s[:4]
	printSlice(s)

	fmt.Println("Drop its first two values. s = s[2:]")
	s = s[2:]
	printSlice(s)

	fmt.Println("setiap drop akan mengurangi kapasitas. s = s[2:]")
	s = s[2:]
	printSlice(s)
	s = s[:2]
	s = s[1:]
	printSlice(s)
}

func printSlice(s []int) {
	// fmt.Println("desimal:")
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// fmt.Println("biner:")
	// fmt.Printf("len=%b cap=%b %v\n", len(s), cap(s), s)
	// fmt.Println("oktal:")
	// fmt.Printf("len=%o cap=%o %v\n", len(s), cap(s), s)
	// fmt.Println("hexa:")
	// fmt.Printf("len=%x cap=%x %v\n", len(s), cap(s), s)
	fmt.Println("==============================")
}
