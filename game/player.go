package game

import "github.com/twmb/algoimpl/go/graph"
import (
    "bufio"
	"math/rand"
	"time"
    "fmt"
    "os"
    "strconv"
    "strings"
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

type HumanPlayer struct {
}

func (player *HumanPlayer) MakeMove(board *Board) int {
    reader := bufio.NewReader(os.Stdin)

    // Show the board and ask for input
    board.Print()
    fmt.Print("Enter column (1-7): ")
    text, _ := reader.ReadString('\n')
    move, _ := strconv.Atoi(strings.TrimSpace(text))
    fmt.Println(move)

    // User move will be 1-indexed. We want 0 indexed
    board.MakeMove(move - 1)

    return move
}

type SmartPlayer struct {
    piece byte
}

func NewSmartPlayer(playerIdx int) *SmartPlayer {
    // Set it to player 0's turn
    return &SmartPlayer{piece: Tokens[playerIdx]}
}

func (player *SmartPlayer) MakeMove(board *Board) int {
    move := 1

    return move
}



func buildMoveTree(board *Board, token byte) (*graph.Graph, *graph.Node) {
    g := graph.New(graph.Directed)
    startNode := g.MakeNode()

    var val interface{} = board.CalcPlayerValue(token)
    startNode.Value = &val

    return g, &startNode
}

