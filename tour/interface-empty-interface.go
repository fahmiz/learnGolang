package main

import "fmt"

func main() {
	explanation()

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = 42.5
	describe(i)

	i = true
	describe(i)

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func explanation() {
	fmt.Println("Explanation:")
	fmt.Println("An empty interface may hold values of any type. (Every type implements at least zero methods.)")
	fmt.Println("Empty interfaces are used by code that handles values of unknown type.")
	fmt.Println("For example, fmt.Print takes any number of arguments of type interface{}.")
	fmt.Println()
}
