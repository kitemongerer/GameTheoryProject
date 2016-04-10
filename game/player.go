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

var Decay = .95

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

    if (move < 1 || move > 7 || !board.IsValidMove(move - 1)) {
        return player.MakeMove(board)
    }
    //fmt.Println(move)

    // User move will be 1-indexed. We want 0 indexed
    board.MakeMove(move - 1)

    return move
}

type SmartPlayer struct {
    Piece byte
    NumLayers int
}

func NewSmartPlayer(playerIdx int, numLayers int) *SmartPlayer {
    return &SmartPlayer{Piece: Tokens[playerIdx], NumLayers: numLayers}
}

func (player *SmartPlayer) MakeMove(board *Board) int {
    g, startNode, _ := buildMoveTree(player.NumLayers, board, player.Piece)

    _, arr := backwardsInduct(g, startNode, player.Piece, board, Decay)

    board.MakeMove(arr[0])
    return arr[0]
}

func backwardsInduct(g *graph.Graph, startNode *graph.Node, token byte, originalBoard *Board, decay float64) (int, []int) {
    if len(g.Neighbors(*startNode)) > 0 {
        // most negative value
        value := -int(^uint(0)  >> 1)
        idx := 0
        //fmt.Printf("Token: %c ", token)
        for i, node := range g.Neighbors(*startNode) {
            tmpVal, _ := backwardsInduct(g, &node, nextToken(token), originalBoard, decay * decay)
          //  fmt.Printf("%d, ", tmpVal)
            if tmpVal > value {
                value = tmpVal
                idx = i
            // If there are two equal values, choose randomly
            } else if tmpVal == value {
                source := rand.NewSource(time.Now().UnixNano())
                rand := rand.New(source)
                
                // Random between 0 and 1
                tmp := rand.Intn(2)
                if tmp == 0 {
                    value = tmpVal
                    idx = i
                }
            }
        }
        //fmt.Printf("\n")
  
        chosenMoveList := (*g.Neighbors(*startNode)[idx].Value).([]int)
        return -value, chosenMoveList
    } else {
        val := buildBoardFromMoveList((*startNode.Value).([]int), originalBoard).CalcPlayerValue(nextToken(token))
        return int(float64(val) * decay), (*startNode.Value).([]int)
    }
}

func nextToken(token byte) byte {
    if Tokens[0] == token {
        return Tokens[1]
    } else {
        return Tokens[0]
    }
}

func buildMoveTree(numLayers int, board *Board, token byte) (*graph.Graph, *graph.Node, *[]graph.Node) {
    g := graph.New(graph.Directed)
    startNode := g.MakeNode()
    valSlice := make([]int, 0, 1)
    *startNode.Value = valSlice

    tmp := make([]graph.Node, 1)
    tmp[0] = startNode
    nodeList := &tmp

    for i:= 0; i < numLayers; i++ {
        newNodeList := make([]graph.Node, 0, len(*nodeList) * NumCols)
        for _, node := range *nodeList {
            nodeBoard := buildBoardFromMoveList((*node.Value).([]int), board)
            
            // Don't keep building out tree if the game is over
            if (!nodeBoard.CheckEndGame()) {
                buildMoveTreeLayer(nodeBoard, g, &node)
                newNodeList = append(newNodeList, g.Neighbors(node)...)
            }
        }
        nodeList = &newNodeList
    }

    return g, &startNode, nodeList
}

func buildMoveTreeLayer(board *Board, g *graph.Graph, startNode *graph.Node) {
    valSlice := (*startNode.Value).([]int)

    for i := 0; i < NumCols; i++ {
        if (board.IsValidMove(i)) {
            newNode := g.MakeNode()
            tmpBoard := *board
            tmpBoard.MakeMove(i)

            // Create new move history array
            tmp := make([]int, len(valSlice), cap(valSlice) + 1)
            copy(tmp, valSlice)
            tmp = append(tmp, i)
            *newNode.Value = tmp

            g.MakeEdge(*startNode, newNode)
        }
    }
}

func buildBoardFromMoveList(moveList []int, oldBoard *Board) *Board {
    board := oldBoard.DuplicateBoard()
    for _, val := range moveList {
        board.MakeMove(val)
    }
    return board
}
