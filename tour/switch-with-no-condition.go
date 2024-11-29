package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Perintah switch tanpa sebuah kondisi sama seperti switch true.")
	fmt.Println("Konstruksi ini merupakan cara yang bersih untuk menulis rantaian if-then-else yang panjang.")
	t := time.Now()
	switch {
	case t.Hour() < 9:
		fmt.Println("Selamat pagi!")
	case t.Hour() < 14:
		fmt.Println("Selamat siang!")
	case t.Hour() < 17:
		fmt.Println("Selamat sore.")
	default:
		fmt.Println("Selamat malam.")
	}
}
