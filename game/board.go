package game

import (
	"fmt"
	"strings"
)

var Tokens = [2]byte{'X', 'O'}

// Represents player values if the section contains a certain
// configuration of tokens (T) and empty spaces
var ConfigValues = map[string]int{
    "T "	:	1,
    " T"	:	1,
    " T "	:	2,
    "TT "	:	3,
    " TT"	:	3,
    " TT "	:	4,
    "TTT "	:	5,
    " TTT"	:	5,
    " TTT "	:	100,
    
    // Maximum integer value since it is a win
    "TTTT"	:	int(^uint(0)  >> 1),
}

const NumCols = 7
const numRows = 6
const numDiags = 12

type Board struct {
	board [NumCols][numRows]byte
	WhoseTurn int
	ValidMoves [NumCols]bool
	Winner byte
}

func NewBoard() *Board {
	var b [NumCols][numRows]byte

	// Initialize all tiles to empty
	for c, _ := range b {
		for r, _ := range b[c] {
			b[c][r] = ' '
		}
	}

	var vm = [NumCols]bool{true, true, true, true, true, true, true}

	// Set it to player 0's turn
	return &Board{WhoseTurn: 0, board: b, ValidMoves: vm, Winner: ' '}
}

func (oldBoard *Board) DuplicateBoard() *Board {
	var b [NumCols][numRows]byte

	// Initialize all tiles to old board tiles
	for c, _ := range b {
		for r, _ := range b[c] {
			b[c][r] = oldBoard.board[c][r]
		}
	}

	var vm [NumCols]bool
	for i := 0; i < NumCols; i++ {
		vm[i] = oldBoard.ValidMoves[i]
	}

	// Set it to player 0's turn
	return &Board{WhoseTurn: oldBoard.WhoseTurn, board: b, ValidMoves: vm, Winner: oldBoard.Winner}
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

// Takes in a function used to calculate the value of a configuration of tokens
// Returns total board valuation
func (board *Board) checkBoardValue(valueFunction func(string) int) int {
	boardValue := 0

	// Check each column
	for _, col := range board.board {
		val := valueFunction(string(col[:numRows]))
		if val == int(^uint(0)  >> 1) || val == -int(^uint(0)  >> 1) {
			return val
		}
		boardValue += val
	}

	// Check each row
	for i := 0; i < numRows; i++ {
		rowSlice := make([]byte, NumCols)

		for j, col := range board.board {
			rowSlice[j] = col[i]
		}

		val := valueFunction(string(rowSlice[:NumCols]))
		if val == int(^uint(0)  >> 1) || val == -int(^uint(0)  >> 1) {
			return val
		}
		boardValue += val
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

		val := valueFunction(string(leftDiagSlice[:len(leftDiagSlice)]))
		if val == int(^uint(0)  >> 1) || val == -int(^uint(0)  >> 1) {
			return val
		}
		boardValue += val

		val = valueFunction(string(rightDiagslice[:len(rightDiagslice)]))
		if val == int(^uint(0)  >> 1) || val == -int(^uint(0)  >> 1) {
			return val
		}
		boardValue += val
	}
	
	return boardValue
}

func (board *Board) CalcPlayerValue(token byte) int {
	// checkSectionValue always returns value for player X
	if token == Tokens[0] {
		return board.checkBoardValue(board.checkSectionValue)	
	} else {
		return -board.checkBoardValue(board.checkSectionValue)
	}
}

func (board *Board) checkSectionValue(s string) int {
	sectionValue := 0

	for tokenIdx, val := range Tokens {
		tokenString := ""

		// Convert 0 to 1 and 1 to - 1
		mul := -1 * (tokenIdx * 2 - 1)

		// Check for win
		if strings.Contains(strings.Replace(s, string(val), "T", -1), "TTTT") {
			return mul * ConfigValues["TTTT"]
		}

		// Find idx of first token
		idx := strings.Index(s, string(val))

		// Make sure token is found
		if idx >= 0 {

			// Check if empty space before token
			if idx != 0 && s[idx - 1] == ' ' {
				tokenString += " "
			}

			// Find all tokens adjacent to this one
			for i := idx; i < len(s); i++ {
				if s[i] == val {
					tokenString += string(val)
				} else {
					// Check for ending space after token string
					if s[i] == ' ' {
						tokenString += " "
					}

					// Replace token with generic token to search map
					tokenString = strings.Replace(tokenString, string(val), "T", -1)
					
					// Don't need to check if key exists because 0 will be returned if it doesn't
					sectionValue += mul * ConfigValues[tokenString]

					// Reset token string and set i to next token
					tokenString = ""
					tmpI := strings.Index(s[i + 1:], string(val))

					// If not found, end loop
					if tmpI < 0 { break }

					i = tmpI + i

					// Check if empty space before token
					if i != 0 && s[i] == ' ' {
						tokenString += " "
					}
				}
			}

			// Add value of last token string
			// Replace token with generic token to search map
			tokenString = strings.Replace(tokenString, string(val), "T", -1)
			
			// Don't need to check if key exists because 0 will be returned if it doesn't
			sectionValue += mul * ConfigValues[tokenString]
		}
	}

	return sectionValue
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
		for c := 0; c < NumCols; c++ {
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

func Reverse(value string) string {
    // Convert string to rune slice.
    // ... This method works on the level of runes, not bytes.
    data := []rune(value)
    result := []rune{}

    // Add runes in reverse order.
    for i := len(data) - 1; i >= 0; i-- {
	result = append(result, data[i])
    }

    // Return new string.
    return string(result)
}
