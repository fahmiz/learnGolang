package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(5, 12)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	kali := func(x, y float64) float64 {
		return x * y
	}
	fmt.Println(kali(10, 10))
	fmt.Println(compute(kali))
}
