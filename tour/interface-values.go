package main

import (
	"fmt"
	"math"
)

type Inter interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println("t.S:", t.S)
}

type F float64

func (f F) M() {
	fmt.Println("f:", f)
}

type Integ int8

func (itg Integ) M() {
	fmt.Println("itg:", itg)
}

func describe(i Inter) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i Inter

	i = &T{"Hello"}
	describe(i)
	i.M()

	fmt.Println()

	i = F(math.Pi)
	describe(i)
	i.M()

	fmt.Println()

	i = Integ(90)
	describe(i)
	i.M()
}
