package game

import (
	"fmt"
	"strings"
)

var Tokens = [2]byte{'X', 'O'}
const numCols = 7
const numRows = 6
const numDiags = 12

type Board struct {
	board [numCols][numRows]byte
	WhoseTurn int
	ValidMoves [numCols]bool
	Winner byte
	GameOver bool
}

func Initialize() *Board {
	var b [numCols][numRows]byte

	// Initialize all tiles to empty
	for c, _ := range b {
		for r, _ := range b[c] {
			b[c][r] = ' '
		}
	}

	var vm = [numCols]bool{true, true, true, true, true, true, true}

	// Set it to player 0's turn
	return &Board{WhoseTurn: 0, board: b, ValidMoves: vm, Winner: ' '}
}

func (board *Board) MakeMove(col int) {
	// Put corresponding piece on board in lowest open spot on column
	for i, val := range board.board[col] {
		if (val == ' ') {
			board.board[col][i] = Tokens[board.WhoseTurn]

			// If column is full, set it as invalid move
			if (i == numRows - 1) {
				board.ValidMoves[col] = false;
			}
			break
		}
	}
	
	board.WhoseTurn = (board.WhoseTurn + 1) % 2
}

func (board *Board) IsValidMove(col int) bool {
	// Check if the move is valid
	return board.ValidMoves[col]
}

func (board *Board) checkForWin() bool {
	// Check each column
	for _, col := range board.board {
		if board.checkSectionWin(string(col[:numRows])) {
			return true
		}
	}

	// Check each row
	for i := 0; i < numRows; i++ {
		rowSlice := make([]byte, numCols)

		for j, col := range board.board {
			rowSlice[j] = col[i]
		}
		
		if board.checkSectionWin(string(rowSlice[:numRows])) {
			return true
		}
	}

	// Check each diagonal
	for i := 0; i < numDiags; i++ {
		var leftDiagSlice, rightDiagslice []byte
		// No diagonal longer than 6
		if (i > 5) {
			leftDiagSlice = make([]byte, 6 - (i + 1) % 7)
			rightDiagslice = make([]byte, 6 - (i + 1) % 7)
		} else {
			leftDiagSlice = make([]byte, i + 1)
			rightDiagslice = make([]byte, i + 1)
		}
		
		for j := 0; j < len(leftDiagSlice); j++ {
			leftDiagSlice[j] = board.board[Min(i, 6) - j][Max(0, i - 6) + j]
			// Same as left but flipped over horizontal axis
			rightDiagslice[j] = board.board[Min(i, 6) - j][5 - (Max(0, i - 6) + j)]
		}

		if board.checkSectionWin(string(leftDiagSlice[:len(leftDiagSlice)])) {
			return true
		}
		if board.checkSectionWin(string(rightDiagslice[:len(rightDiagslice)])) {
			return true
		}
	}
	

	return false
}

func (board *Board) checkSectionWin(s string) bool {
	didWin := false
	if strings.Contains(s, "XXXX") {
		board.Winner = 'X'
		didWin = true
	} else if strings.Contains(s, "OOOO") {
		board.Winner = 'O'
		didWin = true
	}

	return didWin
}

func (board *Board) checkEndGame() {

}

func (board *Board) Print() {
	fmt.Println("\n  1   2   3   4   5   6   7")
	fmt.Println("+---+---+---+---+---+---+---+")

	// Print each row
	for r := numRows - 1; r >= 0; r-- {
		for c := 0; c < numCols; c++ {
			fmt.Printf("| %c ", board.board[c][r])
		}

		fmt.Println("|")
		fmt.Println("+---+---+---+---+---+---+---+")
	}

}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}