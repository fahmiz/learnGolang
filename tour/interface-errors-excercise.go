package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 1
	for i < 30 {
		z -= (z*z - x) / (2 * z)
		i++
	}
	return z
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	res1 := Sqrt(float64(e))
	result := fmt.Sprintf("akar dari %.0f = %.0f", e, res1)
	if e < 0 {
		result = fmt.Sprintf("Error Message: cannot Sqrt negative number: %.0f", e)
	}
	return result
}

func main() {
	fmt.Println("soal latihan: https://go.dev/tour/methods/20")
	fmt.Println(ErrNegativeSqrt(4))
	fmt.Println(ErrNegativeSqrt(-4)) //ini harus error
}
