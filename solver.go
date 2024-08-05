package sudoku

// Solve function tries to solve the Sudoku puzzle.
// It returns true if the board is successfully solved.
func Solve(board []string) bool {
	// Find the first empty cell
	row, col, found := FindEmpty(board)
	if !found {
		// If no empty cell is found, the puzzle is solved
		return true
	}

	// Try placing numbers 1 to 9 in the empty cell
	for num := byte('1'); num <= byte('9'); num++ {
		if IsValid(board, row, col, num) {
			// Place the number if it's valid
			board[row] = board[row][:col] + string(num) + board[row][col+1:]

			// Recursively attempt to solve the board
			if Solve(board) {
				return true
			}

			// If placing the number doesn't lead to a solution, reset the cell
			board[row] = board[row][:col] + "." + board[row][col+1:]
		}
	}

	// Return false if no solution is found
	return false
}

// FindEmpty function finds the first empty cell on the board.
// It returns the row and column of the empty cell, or -1, -1 if no empty cell is found.
func FindEmpty(board []string) (int, int, bool) {
	for i, row := range board {
		for j, cell := range row {
			if cell == '.' {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

// IsValid function checks if placing 'num' at board[row][col] is valid.
func IsValid(board []string, row, col int, num byte) bool {
	// Check the row for the same number
	for _, cell := range board[row] {
		if cell == rune(num) {
			return false
		}
	}

	// Check the column for the same number
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check the 3x3 subgrid for the same number
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}
