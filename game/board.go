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
}

func NewBoard() *Board {
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

func (board *Board) checkBoardValue(valueFunction func(string) int) int {
	boardValue := 0

	// Check each column
	for _, col := range board.board {
		boardValue += valueFunction(string(col[:numRows]))
	}

	// Check each row
	for i := 0; i < numRows; i++ {
		rowSlice := make([]byte, numCols)

		for j, col := range board.board {
			rowSlice[j] = col[i]
		}

		boardValue += valueFunction(string(rowSlice[:numCols]))
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

		boardValue += valueFunction(string(leftDiagSlice[:len(leftDiagSlice)]))
		boardValue += valueFunction(string(rightDiagslice[:len(rightDiagslice)]))
	}
	
	return boardValue
}

func calcPlayerValue(board *Board, token byte) int {
    return 0
}

func (board *Board) checkSectionValue(s string) int {
	return 0
}

func (board *Board) checkForWin() bool {
	return board.checkBoardValue(board.checkSectionWin) != 0
}

func (board *Board) checkSectionWin(s string) int {
	didWin := 0
	if strings.Contains(s, "XXXX") {
		board.Winner = 'X'
		didWin = 1
	} else if strings.Contains(s, "OOOO") {
		board.Winner = 'O'
		didWin = 1
	}

	return didWin
}

func (board *Board) CheckEndGame() bool {
	isBoardFull := true
	for _, col := range board.board {
		for _, val := range col {
			if val == ' ' {
				isBoardFull = false
			}
		}
	}

	// Have to perform CheckForWin in case win occurs on full board
	return board.checkForWin() || isBoardFull
	return false
	
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