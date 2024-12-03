package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Println("sum1:", sum)

	return func(x int) int {
		sum += x
		fmt.Println("x:", x)
		fmt.Println("sum2:", sum)
		return sum
	}
}

func main() {
	pos, neg, pos3 := adder(), adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			"i:", i,
			"pos:", pos(i),
			"neg:", neg(-2*i),
			"pos3:", pos3(3*i-1),
		)
		fmt.Println()
	}
}
