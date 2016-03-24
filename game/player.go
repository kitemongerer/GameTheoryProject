package game

import (
	"math/rand"
	"time"
)

type Player interface {
	MakeMove(board *Board) int
}

type RandomPlayer struct {
}

func (player *RandomPlayer) MakeMove(board *Board) int {
	// Setup rand
	source := rand.NewSource(time.Now().UnixNano())
    rand := rand.New(source)
    
    // Random between 0 and 6
    move := rand.Intn(7)
    for !board.IsValidMove(move) {
    	move = rand.Intn(7)
    }

    board.MakeMove(move)

	return move
}