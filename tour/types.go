package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	st     string     = "this is string variable"
	fl     float32    = 7.5 / 13.7
	bt     byte       = 255
	rn     rune       = -2147483648
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", st, st)
	fmt.Printf("Type: %T Value: %v\n", fl, fl)
	fmt.Printf("Type: %T Value: %v\n", bt, bt)
	fmt.Printf("Type: %T Value: %v\n", rn, rn)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("nilai kosong untuk int, float64, bool, string adalah:")
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var x, y int = 3, 4
	var flo float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(flo)
	fmt.Println("konversi tipe:")
	fmt.Println(x, y, z)

	fmt.Println("Inferensi tipe")
	v1 := 42
	fmt.Printf("v1 bertipe %T\n", v1)
	v2 := 42.5
	fmt.Printf("v2 bertipe %T\n", v2)
	v3 := "Hehehe"
	fmt.Printf("v3 bertipe %T\n", v3)
}
