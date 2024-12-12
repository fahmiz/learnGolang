package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println("s:", s)

	s, ok := i.(string) //a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
	fmt.Println("s, ok:", s, ok)

	f, ok := i.(float64)
	fmt.Println("f, ok:", f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
