package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 1
	for i < 30 {
		//fmt.Print(z, " | ")
		z -= (z*z - x) / (2 * z)
		i++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(100))
}
