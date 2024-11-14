package main

import (
	"fmt"
	"math"
)

func main() {
	// Pada bahasa Go, sebuah nama diekspor jika diawali dengan huruf besar.
	fmt.Println("math.Pi adalah ", math.Pi)
	fmt.Println("2 ^ 3 adalah ", math.Pow(2.0, 3.0))

	// jika fmt.Println(math.pi) maka error
	// tour/exported-names.go:9:14: undefined: math
	// fmt.Println(math.pow(2.0, 3.0)) maka error
}
