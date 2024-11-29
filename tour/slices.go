package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:6]
	fmt.Println("a[bawah : atas]")
	fmt.Println("Notasi di atas memotong rentang dari slice a yang mengikutkan elemen bawah, tapi tidak memasukan bagian terakhir (atas).")
	fmt.Println("sehingga a[bawah] sampai a[atas-1]")
	fmt.Println(s)
}
