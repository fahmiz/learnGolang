package main

import "fmt"

func updateValue(val int) {
	val = 10 // Tidak memengaruhi nilai asli
	fmt.Println("val = 10")
}

func updateValueWithPointer(val *int) {
	*val = 10 // Memengaruhi nilai asli melalui pointer
	fmt.Println("*val = 10")
}

func main() {
	a := 5
	updateValue(a)
	fmt.Println("Output: 5")
	fmt.Println(a)

	updateValueWithPointer(&a)
	fmt.Println("Output: 10")
	fmt.Println(a)
	fmt.Println()

	i, j := 42, 9

	p := &i // menunjuk ke i
	q := i
	fmt.Println("q:", q)   // baca i tidak lewat pointer
	fmt.Println("*p:", *p) // baca i lewat pointer

	q = 21                    // set i tidak lewat pointer
	fmt.Println("q = 21:", q) // baca i lewat pointer
	fmt.Println("i:", i)      // lihat nilai terbaru dari i (ga berubah)

	*p = 21                     // set i lewat pointer
	fmt.Println("*p = 21:", *p) // baca i lewat pointer
	fmt.Println("i:", i)        // lihat nilai terbaru dari i (berubah)
	fmt.Println()

	x := j
	x = x / 3
	fmt.Println("x:", x)
	fmt.Println("j:", j)
	fmt.Println()

	p = &j                 // p menunjuk ke j
	*p = *p / 3            // bagi nilai j lewat pointer
	fmt.Println("*p:", *p) // baca i lewat pointer
	fmt.Println("j:", j)   // lihat nilai terbaru dari j

}
