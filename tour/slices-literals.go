package main

import "fmt"

func main() {
	fmt.Println("Menginisialisasi slice mirip dengan array tapi tanpa mendefinisikan panjangnya.")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("q := []int{2, 3, 5, 7, 11, 13}")
	fmt.Println("q:", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("r := []bool{true, false, true, true, false, true}")
	fmt.Println("r:", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("s:", s)
}
