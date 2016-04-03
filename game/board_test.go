package game

import "testing"

func TestInitialize(t *testing.T) {
	board := NewBoard()

	if board.WhoseTurn != 0 {
		t.Error("Board should initialize to player 0's turn.")
	}

	for _, arr := range board.board {
		for _, val := range arr {
			if val != ' ' {
				t.Error("Board should initialize with a board full of blank spaces.")
				break
			}	
		}
	}
}

func TestMakeMove(t *testing.T) {
	board := NewBoard()

	board.MakeMove(1)

	if board.WhoseTurn != 1 {
		t.Error("After first move, it should be player 1's turn.")
	}

	if board.board[1][0] != 'X' {
		t.Error("Correct square on board should change")
	}

	board.MakeMove(2)

	if board.WhoseTurn != 0 {
		t.Error("After second move, it should be player 0's turn.")
	}

	if board.board[2][0] != 'O' {
		t.Error("Correct square on board should change")
	}

	board.MakeMove(2)

	if board.board[2][1] != 'X' {
		t.Error("Moving on a col with 1 token in it should result in new token being placed above that token")
	}
}

func TestMakeMoveAndValidMove(t *testing.T) {
	board := NewBoard()

	for i:=0; i < 6; i++ {
		board.MakeMove(1)
	}
	
	if board.ValidMoves[1] {
		t.Error("Make move is not setting move to invalid after column is full.")
	}

	if board.IsValidMove(1) {
		t.Error("IsValidMove is somehow returning wrong value.")	
	}
}

func TestSectionWin(t *testing.T) {
	board := NewBoard()

	// Check for X win
	if board.checkSectionWin("XXXX   ") != 1 {
		t.Error("Should win with four X's in a row")	
	}

	if board.Winner != 'X' {
		t.Error("Must set winner after win detected")
	}

	board = NewBoard()

	// Check for O win
	if board.checkSectionWin("  OOOO ") != 1 {
		t.Error("Should win with four O's in a row")	
	}

	if board.Winner != 'O' {
		t.Error("Must set winner after win detected")
	}

	board = NewBoard()

	// Check for not win
	if board.checkSectionWin("OXXX  X") == 1 {
		t.Error("Should not win with only three in a row")	
	}

	if board.Winner != ' ' {
		t.Error("There is no winner")
	}
}

func TestCheckForWin(t *testing.T) {
	board := NewBoard()

	if board.checkForWin() {
		t.Error("No one has moved there shouldn't be a winner")
	}

	// Set up win across
	for i := 6; i >= 3; i-- {
		board.MakeMove(i)
		if (i != 3) {
			board.MakeMove(i)
		}
	}

	if !board.checkForWin() {
		t.Error("Horizontal win: X should be the winner")
	}

	board = NewBoard()

	// Set up win vertical
	for i := 0; i < 4; i++ {
		board.MakeMove(0)
		if (i != 3) {
			board.MakeMove(1)
		}
	}
	
	if !board.checkForWin() {
		t.Error("Vertical win: X should be the winner")
	}

	board = NewBoard()

	// Set up win diagonally
	board.MakeMove(2)
	board.MakeMove(3)
	for i := 0; i < 7; i++ {
		board.MakeMove(i % 4)
		if (i != 3) {
			board.MakeMove((i + 1) % 4)
		}
	}
	
	if !board.checkForWin() {
		t.Error("Diagonal win: X should be the winner")
	}
}

func TestCheckEndGameWin(t *testing.T) {
	board := NewBoard()

	// Set up win across
	for i := 0; i < 4; i++ {
		board.MakeMove(i)
		if (i != 3) {
			board.MakeMove(i)
		}
	}

	if !board.CheckEndGame() {
		t.Error("One player has won. Should be end of game!")
	}
}

func TestCheckEndGameFull(t *testing.T) {
	board := NewBoard()

	// Fill board
	for i := 0; i < numCols * numRows; i++ {
		if i % 2 == 0 {
			board.MakeMove(i % 7)	
		} else {
			board.MakeMove(i % 7)
			board.MakeMove(i % 7)
		}
	}

	// To make there be no winner
	board.board[4][3] = 'X'
	board.board[2][3] = 'Y'
	board.board[5][2] = 'Y'

	if !board.CheckEndGame() {
		t.Error("Board is full. Should be end of game!")
	}

	if board.Winner != ' ' {
		t.Error("No winner. Winner marker should be empty.")
	}

}

func TestCheckEndGameNotFullNotWin(t *testing.T) {
	board := NewBoard()

	if board.CheckEndGame() {
		t.Error("No moves have been made. Game shouldn't be over")
	}
}

func TestCheckSectionValueJustXTokens(t *testing.T) {
	board := NewBoard()

	var sectionValTests = []struct {
	  s        string // input
	  expected int // expected result
	}{
	  {"       ", 0},
	  {"      X", ConfigValues["T "]},
	  {"    X  ", ConfigValues[" T "]},
	  {"     XX", ConfigValues["TT "]},
	  {"    XX ", ConfigValues[" TT "]},
	  {"    XXX", ConfigValues["TTT "]},
	  {"   XXX ", ConfigValues[" TTT "]},
	  {"  XXXX ", ConfigValues["TTTT"]},
	}

	for _, tt := range sectionValTests {
		actual := board.checkSectionValue(tt.s)
		if actual != tt.expected {
			t.Errorf("Section should have correct value [%d] not %d for: \"%s\"", tt.expected, actual, tt.s)
		}
	}
}

func TestCheckSectionValueJustYTokens(t *testing.T) {
	board := NewBoard()

	var sectionValTests = []struct {
	  s        string // input
	  expected int // expected result
	}{
	  {"      O", -ConfigValues["T "]},
	  {"   OOO ", -ConfigValues[" TTT "]},
	  {"  OOOO ", -ConfigValues["TTTT"]},
	}

	for _, tt := range sectionValTests {
		actual := board.checkSectionValue(tt.s)
		if actual != tt.expected {
			t.Errorf("Section should have correct value [%d] not %d for: \"%s\"", tt.expected, actual, tt.s)
		}
	}
}

func TestCheckSectionValueBothTokens(t *testing.T) {
	board := NewBoard()

	var sectionValTests = []struct {
	  s        string // input
	  expected int // expected result
	}{
	  {"      XO", ConfigValues["T "]},
	  {"   XOX ", 2 * ConfigValues["T "]},
	  {" XXOOOX " , ConfigValues["T "] + ConfigValues["TT "]},
	}

	for _, tt := range sectionValTests {
		actual := board.checkSectionValue(tt.s)
		if actual != tt.expected {
			t.Errorf("Section should have correct value [%d] not %d for: \"%s\"", tt.expected, actual, tt.s)
		}
	}
}

func TestCalcPlayerValueEmptyBoard(t *testing.T) {
	board := NewBoard()
	val := board.CalcPlayerValue('X')

	if val != 0 {
		t.Error("Starting Node's value should be 0 on an empty board")
	}
}

