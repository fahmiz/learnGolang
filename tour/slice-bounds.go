package main

import "fmt"

func main() {
	fmt.Println("Saat memotong, anda bisa mengindahkan batas bawah atau atas sehingga Go akan menggunakan nilai default-nya.")
	fmt.Println("var a [10]int = a[0:10] = a[:10] = a[0:] = a[:]")
	fmt.Println()

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s := []int{2, 3, 5, 7, 11, 13}")
	fmt.Println("s =", s[:])

	s = s[1:4]
	fmt.Println("s = s[1:4] :", s)

	s = s[:2]
	fmt.Println("s = s[:2] :", s)

	s = s[1:]
	fmt.Println("s = s[1:] :", s)
}
