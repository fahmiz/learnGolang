package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim
}

func pow3(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// v tidak dapat digunakan disini
	return lim
}

func main() {
	fmt.Println("sqrt(-25) =", sqrt(-25))
	fmt.Println("sqrt(4) =", sqrt(4))
	fmt.Println()

	fmt.Println("pow2(3, 2, 8) =", pow2(3, 2, 8))
	fmt.Println("pow2(3, 3, 20) =", pow2(3, 3, 20))

	fmt.Println("pow(3, 2, 8) =", pow(3, 2, 8))
	fmt.Println("pow(3, 3, 20) =", pow(3, 3, 20))

	fmt.Println("pow3(3, 2, 10)", pow3(3, 2, 10))
	fmt.Println("pow3(3, 3, 10)", pow3(3, 3, 20))
}
