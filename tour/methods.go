package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
	A, B string
}

// as methods (using receiver)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) mins() float64 {
	return v.X - v.Y
}

func (v Vertex) concatFloat() string {
	x := int(v.X)
	y := int(v.Y)
	sx := fmt.Sprintf("%d", x)
	sy := fmt.Sprintf("%d", y)
	gabung := sx + sy
	return gabung
}

func (v Vertex) concat() string {
	gabung := v.A + v.B
	return gabung
}

// as regular function
func Abs_func(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Str string

// receiver cannot > 1, parameter may > 1
func (a Str) concats(b, c Str) string {
	return string(a) + string(b) + string(c)
}

type MyFloat float64

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	fmt.Println("Method pada golang memungkinkan golang menyerupai oop")
	v := Vertex{3, 4, "aa", "bb"}
	fmt.Println(v.Abs())
	fmt.Println(Abs_func(v)) //as regular function (no receiver)
	fmt.Println(v.mins())
	fmt.Println(v.concatFloat())
	fmt.Println(v.concat())

	var a Str = "12"
	var b Str = "13"
	var c Str = "14"
	fmt.Println(a.concats(b, c))

	f := MyFloat(-math.Sqrt2) //math.Sqrt2 adalah konstanta akar 2 lalu di minus kan
	fmt.Println(f.Abs2())
	f = MyFloat(math.Sqrt2) //math.Sqrt2 adalah konstanta akar 2
	fmt.Println(f.Abs2())
	fmt.Println()

	fmt.Println("Macam-macam konstanta bawaan math:")
	fmt.Println("Pi:", math.Pi)
	fmt.Println("E:", math.E)
	fmt.Println("Phi (Golden Ratio):", math.Phi)
	fmt.Println("Sqrt2:", math.Sqrt2)
	fmt.Println("Ln2 (Natural Log of 2):", math.Ln2)
	fmt.Println("Log10E (Log base 10 of e):", math.Log10E)
}
