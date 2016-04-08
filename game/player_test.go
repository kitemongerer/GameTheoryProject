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
	player := NewSmartPlayer(1, 1)

	if player.Piece != 'O' {
		t.Error("Player 2 should have piece O.")
	}

	player = NewSmartPlayer(0, 1)
	
	if player.Piece != 'X' {
		t.Error("Player 1 should have piece X.")
	}	
}

func TestBuildMoveTree(t *testing.T) {
	// Check 1 deep empty board
	board := NewBoard()
	g, start, _ := buildMoveTree(1, board, 'X')

	if len((*start.Value).([]int)) != 0 {
		t.Error("Starting Node's move history should have length 0 on an empty board")
	}

	if len(g.Neighbors(*start)) != NumCols {
		t.Error("Move tree should contain all columns on empty board")	
	}

	// Check 2 deep empty board
	board = NewBoard()
	g, start, _ = buildMoveTree(2, board, 'X')

	if len((*start.Value).([]int)) != 0 {
		t.Error("Starting Node's move history should have length 0 on an empty board")
	}

	neighbors := g.Neighbors(*start);

	if len(neighbors) != NumCols {
		t.Error("Move tree should contain all columns on empty board")	
	}

	for _, node := range neighbors {
		secondNeighbors := g.Neighbors(node)
		if len(secondNeighbors) != NumCols {
			t.Error("Second layer neighbors should also contain all columns")	
		}
	}

	// Check 1 deep row 1 full
	board = NewBoard()
	for i := 0; i < 6; i++ {
		board.MakeMove(1)
	}

	g, start, _ = buildMoveTree(1, board, 'X')

	if len((*start.Value).([]int)) != 0 {
		t.Error("Starting Node's move history should have length 0 on an empty board")
	}

	if len(g.Neighbors(*start)) != NumCols - 1 {
		t.Error("Move tree should contain all columns except column 1 since it is full")	
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
	board := buildBoardFromMoveList(moveList, NewBoard())

	// Check that no tokens were added
	if board.CalcPlayerValue('X') != 0 {
		t.Error("No moves should have been made from empty move list.")
	}

	moveList = append(moveList, 0)
	board = buildBoardFromMoveList(moveList, NewBoard())
	// Check that token was added
	if board.board[0][0] != 'X' {
		t.Error("Should have built board with single move.")
	}

	moveList = append(moveList, 0)
	board = buildBoardFromMoveList(moveList, NewBoard())
	// Check that first token was added
	if board.board[0][0] != 'X' {
		t.Error("Should have built board with first token in correct spot.")
	}
	// Check that second token was added
	if board.board[0][1] != 'O' {
		t.Error("Should have built board with alternating tokens.")
	}
}

func TestSmartMakeMove(t *testing.T) {
	board := NewBoard()

	board.board[0] = [6]byte{'X', 'X', ' ', ' ', ' ', ' '}
	board.board[1] = [6]byte{'X', 'X', 'O', 'O', ' ', ' '}
	board.board[2] = [6]byte{'X', 'O', 'X', 'X', ' ', ' '}
	board.board[3] = [6]byte{'O', 'X', 'O', 'O', 'O', 'X'}
	board.board[4] = [6]byte{'X', 'O', 'O', 'X', 'O', 'O'}
	board.board[5] = [6]byte{'X', 'X', 'O', 'X', 'O', ' '}
	board.board[6] = [6]byte{' ', ' ', ' ', ' ', ' ', ' '}

	board.WhoseTurn = 1
	//board.Print()
	var p2 = NewSmartPlayer(1, 4)
	move := p2.MakeMove(board)

	//board.Print()
	if move != 2 {
		t.Error("Should have made move to win game.")
	}
}



