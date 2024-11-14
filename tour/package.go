package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Bilangan kesukaan saya adalah", rand.Intn(10))
	fmt.Println("Bilangan kesukaan dia adalah", rand.Intn(100))
	fmt.Println("Bilangan kesukaan mereka adalah", rand.Intn(1000))
}
