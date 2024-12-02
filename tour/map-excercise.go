package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	// Buat map kosong untuk menyimpan hasil
	wCount := make(map[string]int)
	fmt.Println(wCount)

	// Pecah string menjadi slice kata-kata
	fmt.Println(s)
	words := strings.Fields(s)
	fmt.Println(words)

	// Iterasi setiap kata
	for _, word := range words {
		// Tambahkan jumlah kata ke dalam map
		wCount[word]++
		fmt.Println(wCount)
	}

	return wCount
}

func main() {
	// Tes fungsi WordCount
	fmt.Println(WordCount("hello    world hello apa apa apanya dong"))
}
