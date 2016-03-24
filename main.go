package main

import (
	"FinalProject/game"
	"os"
    "os/exec"
    "time"
    "fmt"
)

func main() {
	var b = game.Initialize()
	
	//var p1, p2 game.RandomPlayer

	var p1 game.HumanPlayer
	var p2 game.RandomPlayer

	for !b.CheckEndGame() {
		p1.MakeMove(b)

		// Delay
		duration := time.Second
  		time.Sleep(duration)

		// If p1 didn't win, p2 gets to move
		if (!b.CheckEndGame()) {
			p2.MakeMove(b)
		}

		b.Print()

		// Clear terminal
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
	}

	if b.Winner != ' ' {
		fmt.Printf("\nPlayer %c is the winner!\n\n", b.Winner)
	} else {
		fmt.Println("\nTie game!\n")
	}

	b.Print()
}
