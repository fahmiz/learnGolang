package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	fmt.Println("A map maps keys to values.")
	fmt.Println("The zero value of a map is nil. A nil map has no keys, nor can keys be added.")
	fmt.Println("The make function returns a map of the given type, initialized and ready for use.")
	fmt.Println("m:", m)

	m = make(map[string]Vertex)
	fmt.Println("make: Fungsi bawaan Go yang digunakan untuk mengalokasikan dan menginisialisasi tipe data seperti map, slice, atau channel.")
	fmt.Println("m = make(map[string]Vertex):", m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Bell"] = Vertex{
		40.68432, -74.39966,
	}
	fmt.Println("Bell:", m["Bell"])
	fmt.Println("Bell Labs:", m["Bell Labs"])
	fmt.Println("m:", m)

}
