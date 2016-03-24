package game

import (
	"fmt"
	"strings"
)

var Tokens = [2]byte{'X', 'O'}
const numCols = 7
const numRows = 6

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
	for i, val := range board.board {
		
		if board.checkSectionWin(string(val[:numRows])) {
			return true
		}

		slice := make([]byte, numRows)
		for i, char := range board.board[i] {
			slice[i] = char
		}
		fmt.Println(string(slice[:numRows]))
		if board.checkSectionWin(string(slice[:numRows])) {
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