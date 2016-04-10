package main

import (
	"FinalProject/game"
	"os"
    "os/exec"
    "fmt"
    "strconv"
    "runtime"
    "sync"
    "time"
    "net/http"
    "log"
)

const (
    DEFAULT_PORT = "9000"
)
  
func HelloServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello, World!\n")
}
  
func main() {
    var port string
    if port = os.Getenv("PORT"); len(port) == 0 {
        log.Printf("Warning, PORT not set. Defaulting to %+v\n", DEFAULT_PORT)
        port = DEFAULT_PORT
    }
  
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Printf("ListenAndServe: ", err)
    }


	runtime.GOMAXPROCS(runtime.NumCPU())

	// Get command line input
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		numReps, e := strconv.Atoi(argsWithoutProg[0])
		if e != nil {
			numReps = 1
		}

		wg := &sync.WaitGroup{}

		var victorySlice = make([]byte, numReps)
		var wins = [3]int{0, 0, 0}

		start := time.Now()
		for i := 0; i < numReps; i++ {
			wg.Add(1)
			go executeGame(i, &victorySlice, wg)
		}

		wg.Wait()

		elapsed := time.Since(start)
	    fmt.Println(elapsed)

		for _, winner := range victorySlice {

			if winner == ' ' {
				wins[2]++
			} else if winner == 'X' {
				wins[0]++
			} else if winner == 'O' {
				wins[1]++
			}	
		}

		fmt.Printf("Player 1 won %d times; Player 2 won %d times; There were %d ties.\n", wins[0], wins[1], wins[2])
	} else {
		executeHumanGame()
	}

	fmt.Println("Simulation complete.")
}

func executeGame(idx int, victorySlice *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	var b = game.NewBoard()

	// Switch off who goes first
	b.WhoseTurn = idx % 2

	var p1 = game.NewSmartPlayer(0, 3)
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

func executeHumanGame() {
	var b = game.NewBoard()

	var p1 game.HumanPlayer
	var p2 = game.NewSmartPlayer(1, 5)

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
