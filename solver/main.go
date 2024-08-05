package main

import (
	"fmt"

	"sudoku"
)

func main() {
	board := []string{
		".96.4...1",
		"1...6...4",
		"5.481.39.",
		"..795..43",
		".3..8....",
		"4.5.23.18",
		".1.63..59",
		".59.7.83.",
		"..359...7",
	}
	sudoku.Solve(board)

	printBoard(board)

}

func printBoard(board []string) {
	for _, row := range board {
		fmt.Println(row)
	}
}
