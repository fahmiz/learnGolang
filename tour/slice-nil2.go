package main

import (
	"fmt"
)

func fetchData(condition bool) []int {
	if !condition {
		return nil // Tidak ada data
	}
	return []int{1, 2, 3} // Data tersedia
}

func main() {
	fmt.Println("data1:")
	data := fetchData(false)
	if data == nil {
		fmt.Println("No data found")
		fmt.Println(data)
		fmt.Println("append data")
		data = append(data, 7, 8, 9)
		fmt.Println(data)
		fmt.Println("menghapus data")
		data = nil
		fmt.Println(data)

		fmt.Println()
	}
	fmt.Println("data2:")
	data2 := fetchData(true)
	if data2 == nil {
		fmt.Println("No data found")
	}
	fmt.Println(data2)
}
