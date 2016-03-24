package main

import (
	"FinalProject/game"
)

func main() {
	var b = game.Initialize()
	
	var p1, p2 game.RandomPlayer

	for !b.CheckEndGame() {
		p1.MakeMove(b)

		// If p1 didn't win, p2 gets to move
		if (!b.CheckEndGame()) {
			p2.MakeMove(b)
		}
	}

	b.Print()
}
