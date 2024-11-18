package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Println("j", j)
		}
		fmt.Println()
	}
	fmt.Println("sum", sum)

	sum2 := 1
	for sum2 < 1000 {
		fmt.Println("sum2", sum2)
		sum2 += sum2
	}
	fmt.Println("sum2", sum2)

	// invinite loop
	// for {}

}
