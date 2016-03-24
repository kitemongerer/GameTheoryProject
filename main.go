package main

import (
	"FinalProject/game"
)

func main() {
	var b = game.Initialize()
	
	b.Print()

	b.MakeMove(1)
	b.MakeMove(2)
	b.MakeMove(1)
	b.MakeMove(1)
	b.MakeMove(1)
	b.MakeMove(1)
	b.MakeMove(3)

	b.Print()
}
