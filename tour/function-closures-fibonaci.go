package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int. 0 1 1 2 3 5 8 13 21 34 55 89
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 24; i++ {
		fmt.Print(f(), " ~ ")
	}
}
