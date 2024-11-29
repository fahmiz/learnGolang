package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Slice bisa mengandung tipe apapun, termasuk slice lainnya.")
	// Buat papan tic-tac-toe
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println("board :", board)

	// Giliran untuk para pemain
	board[0][0] = "X"
	fmt.Println("board[0][0]=x:", board)
	printmatrix(board)

	board[2][2] = "O"
	fmt.Println("board[2][2]=0:", board)
	printmatrix(board)

	board[1][2] = "X"
	fmt.Println("board[1][2]=x:", board)
	printmatrix(board)

	board[1][0] = "O"
	fmt.Println("board[1][0]=0:", board)
	printmatrix(board)

	board[0][2] = "X"
	fmt.Println("board[0][2]=x:", board)
	printmatrix(board)

}

func printmatrix(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
