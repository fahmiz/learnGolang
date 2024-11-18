package main

import "fmt"

const (
	Pi = 3.14
	// Buat bilangan yang besar dengan men-shift 1 bit ke kiri 100 kali.
	// Dengan kata lain, bilangan binari 1 diikuti dengan 100 angka nol.
	Big = 1 << 100
	// Shift kembali ke kanan sebanyak 99 kali, sehingga akhirnya menjadi
	// 1<<1, atau 2
	Small = Big >> 99

	Medium     = 7 << 1
	xtra_small = Medium >> 1
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println()

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println()

	fmt.Println(int(Medium))
	biner := fmt.Sprintf("%b", Medium)
	fmt.Println(biner)
	biner2 := fmt.Sprintf("%b", xtra_small)
	fmt.Println(biner2)
}
