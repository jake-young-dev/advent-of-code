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
	OLD_SPOT     = 46
	DOT_SPOT     = 46
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
	var starting point
	for i, l := range filelines {
		board[i] = make([]rune, len(l))
		for ii, c := range l {
			if c == FACING_UP {
				guard = point{
					x:      i,
					y:      ii,
					facing: FACING_UP,
				}
				starting = point{
					x: i,
					y: ii,
				}
			}
			board[i][ii] = c
		}
	}

	lastx := -1
	lasty := -1
	cycleCounter := 0
	log.Println("adding new blockers")
	for i, x := range board {
		log.Println("--")
		log.Printf("trying blockers on row: %d\n", i+1)
		rt := time.Now()
		for j, y := range x {
			if i == starting.x && j == starting.y {
				continue
			}
			if y == BLOCKER {
				continue
			}
			if lastx != -1 && lasty != -1 {
				board[lastx][lasty] = DOT_SPOT
			}
			board[i][j] = BLOCKER
			lastx = i
			lasty = j
			err = nil
			oldMoves := make(map[point]int)

			ecache := -1
			var moved bool
			for err == nil {
				// 	printBoard(board) //uncomment to debug, stepping thru each move
				// 	log.Println("press enter for next move")
				// 	bufio.NewReader(os.Stdin).ReadString('\n')
				// 	board, guard, moved, err = move(board, guard, oldMoves)
				// 	log.Println()
				// 	log.Println("---------------------------------------------")
				board, guard, moved, err = move(board, guard, oldMoves)

				eNum := 0
				for range oldMoves {
					eNum += 1
				}

				if moved && (eNum == ecache) {
					log.Printf("found a cycle with a new blocker on: %d, %d", i, j)
					cycleCounter++
					break
				}

				if moved {
					ecache = eNum
				}
			}

			board[guard.x][guard.y] = DOT_SPOT
			guard.x = starting.x
			guard.y = starting.y
			guard.facing = FACING_UP
		}
		log.Printf("row %d processed in %s, %d total cycles found\n", i+1, time.Since(rt).String(), cycleCounter)
	}

	log.Printf("%d cycles found\n", cycleCounter)
	log.Printf("Execution time: %s", time.Since(st).String())
}

func move(board [][]rune, guard point, old map[point]int) ([][]rune, point, bool, error) {
	//grab new spot
	movedx := guard.x + moveMap[guard.facing].x
	movedy := guard.y + moveMap[guard.facing].y

	//does new spot exist on board
	if len(board) <= movedx || movedx < 0 {
		old[guard] = 0
		return board, guard, false, errors.New("off board move")
	}
	if len(board[movedx]) <= movedy || movedy < 0 {
		old[guard] = 0
		return board, guard, false, errors.New("off board move")
	}

	//is new spot a blocker
	if board[movedx][movedy] == BLOCKER {
		//turn right
		guard.facing = turnMap[guard.facing]
		board[guard.x][guard.y] = guard.facing
		return board, guard, false, nil
	} else {
		//move forward
		board[guard.x][guard.y] = OLD_SPOT
		old[guard] += 1
		board[movedx][movedy] = guard.facing
		guard.x = movedx
		guard.y = movedy
		return board, guard, true, nil
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

//Execution time: 42m2.251345s
