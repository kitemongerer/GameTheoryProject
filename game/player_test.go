package game

import "testing"
import "github.com/twmb/algoimpl/go/graph"

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

func TestBuildMoveTree(t *testing.T) {
	board := NewBoard()
	g, start := buildMoveTree(1, board, 'X')

	if len((*start.Value).([]int)) != 0 {
		t.Error("Starting Node's move history should have length 0 on an empty board")
	}

	if len(g.Neighbors(*start)) != NumCols {
		t.Error("Move tree should contain all columns on empty board")	
	}
}


func TestBuildMoveTreeLayerEmptyBoard(t *testing.T) {
	board := NewBoard()
	g := graph.New(graph.Directed)
    startNode := g.MakeNode()
    valSlice := make([]int, 0, 1)
    *startNode.Value = valSlice
	
	buildMoveTreeLayer(board, g, &startNode)

	if len(g.Neighbors(startNode)) != NumCols {
		t.Error("Move tree should contain all columns on empty board")	
	}

	for i, node := range g.Neighbors(startNode) {
		val := (*node.Value).([]int)

		if len(val) != 1 {
			t.Error("Node's move history should have length 1 after first set of moves")
		}

		if val[0] != i {
			t.Errorf("Node's move history should contain correct moves(%d not %d)", i, val[0])
		}
	}
}

func TestBuildBoardFromMoveList(t *testing.T) {
	moveList := make([]int, 0, 5)
	board := buildBoardFromMoveList(moveList)

	// Check that no tokens were added
	if board.CalcPlayerValue('X') != 0 {
		t.Error("No moves should have been made from empty move list.")
	}

	moveList = append(moveList, 0)
	board = buildBoardFromMoveList(moveList)
	// Check that token was added
	if board.board[0][0] != 'X' {
		t.Error("Should have built board with single move.")
	}

	moveList = append(moveList, 0)
	board = buildBoardFromMoveList(moveList)
	// Check that first token was added
	if board.board[0][0] != 'X' {
		t.Error("Should have built board with first token in correct spot.")
	}
	// Check that second token was added
	if board.board[0][1] != 'O' {
		t.Error("Should have built board with alternating tokens.")
	}
}


