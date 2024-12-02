package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
	"bla": Vertex{
		12, 13,
	},
	"ble": {
		14, 15,
	},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google"])
	fmt.Println(m["bla"])
	fmt.Println(m["ble"]) //If the top-level type is just a type name, you can omit it from the elements of the literal.
	fmt.Println(m["blx"])
}
