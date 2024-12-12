package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v (%T)\n", v, v*2, v)
	case string:
		fmt.Printf("%q is %v bytes long (%T)\n", v, len(v), v)
	default:
		fmt.Printf("I don't know about type (%T)!\n", v)
	}
}

func main() {
	do(21)
	do(21.5)
	do("hello")
	do(true)
}
