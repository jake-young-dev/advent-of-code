package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	FACING_DOWN  = 118
	FACING_LEFT  = 60
	FACING_RIGHT = 62
	FACING_UP    = 94
	BLOCKER      = 35
	OLD_SPOT     = 88
)

type point struct {
	x      int
	y      int
	facing rune
}

var turnMap = map[rune]rune{
	FACING_DOWN:  FACING_LEFT,
	FACING_LEFT:  FACING_UP,
	FACING_UP:    FACING_RIGHT,
	FACING_RIGHT: FACING_DOWN,
}

var moveMap = map[rune]point{
	FACING_LEFT:  {x: 0, y: -1},
	FACING_DOWN:  {x: 1, y: 0},
	FACING_RIGHT: {x: 0, y: 1},
	FACING_UP:    {x: -1, y: 0},
}

func main() {
	st := time.Now()
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	datastr := string(fileBytes)

	filelines := strings.Split(datastr, "\n")

	board := make([][]rune, len(filelines))
	var guard point
	for i, l := range filelines {
		board[i] = make([]rune, len(l))
		for ii, c := range l {
			if c == FACING_UP {
				guard = point{
					x:      i,
					y:      ii,
					facing: FACING_UP,
				}
			}
			board[i][ii] = c
		}
	}

	oldMoves := make(map[[2]int]int)
	err = nil
	for err == nil {
		// printBoard(board)
		// log.Println("press enter for next move")
		// bufio.NewReader(os.Stdin).ReadString('\n')
		board, guard, err = move(board, guard, oldMoves)
		// log.Println()
		// log.Println("---------------------------------------------")
	}

	log.Printf("%d total moves\n", len(oldMoves))
	log.Printf("Execution time: %s", time.Since(st).String())
}

func move(board [][]rune, guard point, old map[[2]int]int) ([][]rune, point, error) {
	//grab new spot
	movedx := guard.x + moveMap[guard.facing].x
	movedy := guard.y + moveMap[guard.facing].y

	//does new spot exist on board
	if len(board) <= movedx || movedx < 0 {
		old[[2]int{guard.x, guard.y}] = 1
		return board, guard, errors.New("off board move")
	}
	if len(board[movedx]) <= movedy || movedy < 0 {
		old[[2]int{guard.x, guard.y}] = 1
		return board, guard, errors.New("off board move")
	}

	//is new spot a blocker
	if board[movedx][movedy] == BLOCKER {
		//turn right
		guard.facing = turnMap[guard.facing]
		board[guard.x][guard.y] = guard.facing
		return board, guard, nil
	} else {
		//move forward
		board[guard.x][guard.y] = OLD_SPOT
		old[[2]int{guard.x, guard.y}] = 1
		board[movedx][movedy] = guard.facing
		guard.x = movedx
		guard.y = movedy
		return board, guard, nil
	}
}

func printBoard(board [][]rune) {
	for _, b := range board {
		for _, c := range b {
			fmt.Printf(string(c))
		}
		fmt.Printf("\n")
	}
}

//Execution time: 1.0196ms
