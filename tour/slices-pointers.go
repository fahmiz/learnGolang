package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("names =", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("a = names[0:2]", "b = names[1:3] =", a, b)
	fmt.Println()

	b[0] = "XXX"
	fmt.Println("b[0] = 'XXX'")
	fmt.Println("a, b =", a, b)
	fmt.Println("names =", names)

}
