package main

import (
	"FinalProject/game"
	"os"
    "fmt"
    "strconv"
    "runtime"
    "sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Get command line input
	argsWithoutProg := os.Args[1:]
	numReps, e := strconv.Atoi(argsWithoutProg[0])
	if e != nil {
		numReps = 1
	}

	wg := &sync.WaitGroup{}

	var victorySlice = make([]byte, numReps)
	var wins = [3]int{0, 0, 0}

	for i := 0; i < numReps; i++ {
		wg.Add(1)
		go executeGame(i, &victorySlice, wg)
	}

	wg.Wait()

	for _, winner := range victorySlice {

		if winner == ' ' {
			wins[2]++
		} else if winner == 'X' {
			wins[0]++
		} else if winner == 'O' {
			wins[1]++
		}	
	}

	fmt.Printf("X won %d times; O won %d times; There were %d ties.\n", wins[0], wins[1], wins[2])

	fmt.Println("Simulation complete.")
}

func executeGame(idx int, victorySlice *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	var b = game.NewBoard()

	// Switch off who goes first
	b.WhoseTurn = idx % 2

	var p1 = game.NewSmartPlayer(0, 1)
	var p2 = game.NewSmartPlayer(1, 1)

	for !b.CheckEndGame() {
		p1.MakeMove(b)
		//b.Print()
		// If p1 didn't win, p2 gets to move
		if (!b.CheckEndGame()) {
			p2.MakeMove(b)
		}
		//b.Print()
		//fmt.Println(b.CalcPlayerValue('O'))
	}

	(*victorySlice)[idx] = b.Winner
}
