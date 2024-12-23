// part of generics
package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			fmt.Printf("index ke %v sama, v = %v, x = %v,s = %v, return=", i, v, x, s)
			return i
		}
		fmt.Printf("index ke %v tidak sama\n", i)
	}
	fmt.Print("semua index tidak sama! return=")
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz", "hello"}
	fmt.Println(Index(ss, "baz"))  //hello is on index 3
	fmt.Println(Index(ss, "helo")) // -1 because not found
	fmt.Println(Index(si, -10))    // -10 is on index 3
}
