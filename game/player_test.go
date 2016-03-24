package game

import "testing"

func TestRandomMakeMove(t *testing.T) {
	var player RandomPlayer
	var moves [7]int

	for i := 0; i < 700; i++ {
		board := Initialize()
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
	board := Initialize()

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