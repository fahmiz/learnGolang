package main

import "fmt"

type Itf interface {
	M()
}

type T struct {
	S string //khusus string
	// S interface{} // bisa semua tipe
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

type Integ8 int8

func (itg Integ8) M() {
	fmt.Println(itg)
}

func main() {
	var i Itf = T{"hello"}
	i.M()
	// i = T{20}
	// i.M()
	// i = T{20.5}
	// i.M()
	i = Integ8(80)
	i.M()
}
