package game

import "testing"

func TestRandomMakeMove(t *testing.T) {
	var player RandomPlayer
	var moves [7]int

	for i := 0; i < 700; i++ {
		board := NewBoard()
		move := player.MakeMove(board)

		if (move > 6 || move < 0) {
			t.Error("Move is out of bounds!")
			break
		}

		if board.board[move][0] == ' ' {
			t.Error("Player should fill in move that they made.")
			break
		}

		moves[move]++
	}

	for _, val := range moves {
		if val < 80 || val > 120 {
			t.Error("Random player's moves don't appear to be random: %v", moves)
		}
	}
}

func TestRandomMakeInvalidMove(t *testing.T) {
	var player RandomPlayer
	board := NewBoard()

	// Fill board except last column
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			board.MakeMove(i)
		}
	}

	for i := 0; i < 6; i++ {
		move := player.MakeMove(board)

		if move != 6 {
			t.Error("Random player trying to make a move in full column.")
		}
	}
}

func TestSmartPlayerInitialize(t *testing.T) {
	player := NewSmartPlayer(1)

	if player.piece != 'O' {
		t.Error("Player 2 should have piece O.")
	}

	player = NewSmartPlayer(0)
	
	if player.piece != 'X' {
		t.Error("Player 1 should have piece X.")
	}	
}


func TestBuildMoveTreeEmptyBoard(t *testing.T) {
	board := NewBoard()
	g, start := buildMoveTree(board, 'X')
	
	if (*start.Value).(int) != 0 {
		t.Error("Starting Node's value should be 0 on an empty board")
	}

	if len(g.Neighbors(*start)) != NumCols {
		t.Error("Move tree should contain all columns on empty board")	
	}

	for i, node := range g.Neighbors(*start) {
		val := *node.Value
		if i == 0 || i == NumCols - 1 {
			if (val).(int) != 3 * ConfigValues["T "] {
				t.Errorf("Player value after moving on end columns should be %d not %d",  3 * ConfigValues["T "], val)	
			}
		} else {
			if (val).(int) != ConfigValues[" T "] + 3 * ConfigValues["T "] {
				t.Errorf("Player value after moving on end columns should be %d not %d", ConfigValues[" T "] + 2 * ConfigValues["T "], val)
			}
		}
	}
}


