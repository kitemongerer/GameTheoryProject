package game

import "testing"

func TestInitialize(t *testing.T) {
	board := Initialize()

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
	board := Initialize()

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
	board := Initialize()

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
	board := Initialize()

	// Check for X win
	if !board.checkSectionWin("XXXX   ") {
		t.Error("Should win with four X's in a row")	
	}

	if board.Winner != 'X' {
		t.Error("Must set winner after win detected")
	}

	board = Initialize()

	// Check for O win
	if !board.checkSectionWin("  OOOO ") {
		t.Error("Should win with four O's in a row")	
	}

	if board.Winner != 'O' {
		t.Error("Must set winner after win detected")
	}

	board = Initialize()

	// Check for not win
	if board.checkSectionWin("OXXX  X") {
		t.Error("Should not win with only three in a row")	
	}

	if board.Winner != ' ' {
		t.Error("There is no winner")
	}
}

func TestCheckForWin(t *testing.T) {
	board := Initialize()

	if board.checkForWin() {
		t.Error("No one has moved there shouldn't be a winner")
	}

	// Set up win across
	for i := 0; i < 4; i++ {
		board.MakeMove(i)
		if (i != 3) {
			board.MakeMove(i)
		}
	}
	board.Print()
	if !board.checkForWin() {
		t.Error("Horizontal win: X should be the winner")
	}

	board = Initialize()

	// Set up win vertical
	for i := 0; i < 4; i++ {
		board.MakeMove(0)
		if (i != 3) {
			board.MakeMove(1)
		}
	}
	board.Print()
	if !board.checkForWin() {
		t.Error("Vertical win: X should be the winner")
	}

	board = Initialize()

	// Set up win diagonally
	board.MakeMove(2)
	board.MakeMove(3)
	for i := 0; i < 7; i++ {
		board.MakeMove(i % 4)
		if (i != 3) {
			board.MakeMove((i + 1) % 4)
		}
	}
	board.Print()
	if !board.checkForWin() {
		t.Error("Diagonal win: X should be the winner")
	}
}

