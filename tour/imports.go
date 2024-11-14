package main

import (
	"fmt"
	"math"
)

func main() {
	const name, age, kwadrat = "Kim", 22, 49
	fmt.Printf("Sekarang anda memiliki %1f masalah.\n", math.Sqrt(kwadrat))
	fmt.Println("Nama :", name)
	fmt.Println("Umur :", age)
	fmt.Println(name, "is", age, "years old")
}
