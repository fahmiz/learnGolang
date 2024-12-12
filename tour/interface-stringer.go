package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// implementasi stringer di fungsi berikut
func (p Person) String() string {
	// fmt.Println("calling stringer Person:")
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

type Person2 struct {
	Name string
	Age  int
}

func (p2 Person2) String() string {
	return fmt.Sprintf("Nama: %v (Usia: %v tahun)\n", p2.Name, p2.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	b := Person{"Bomber Dent", 43}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z, b, b, b, b, b)
	x := Person2{"Arthur Dent", 42}
	y := Person2{"Bomber Dent", 43}
	c := Person2{"Zaphod Beeblebrox", 9001}
	fmt.Println(x, y, c, c, c, c, c)
}
