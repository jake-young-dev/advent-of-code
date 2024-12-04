package main

import (
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

const (
	//equivalent rune's
	M_CHAR = 77
	A_CHAR = 65
	S_CHAR = 83
)

// map characters to next charater in XMAS
var NEXT_TOKEN_VALUE = map[rune]rune{
	M_CHAR: A_CHAR,
	A_CHAR: S_CHAR,
	S_CHAR: 0,
}

// all surrounding spots in matrix for an "X"
var SURROUNDINGS = [][2]int{
	{-1, 1},
	{-1, -1},
	{1, 1},
	{1, -1},
}

// struct to represent each character
type xmastoken struct {
	x    int
	y    int
	xdir int
	ydir int
	next rune
}

func main() {
	st := time.Now()

	//stage 1 - parse data to runes and collect "X" locations

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	dataStr := string(data)
	lines := strings.Split(dataStr, "\r\n")

	lineNumbers := len(lines)
	lineLength := len(lines[0])

	matrix := make([][]rune, lineNumbers)
	//starting locations A.K.A. "M"'s
	var startlocs []xmastoken
	for idx, i := range lines {
		matrix[idx] = make([]rune, lineLength)
		for ndx, j := range i {
			matrix[idx][ndx] = j
			if j == M_CHAR { //find all "starting points"
				startlocs = append(startlocs, xmastoken{
					x:    idx,
					y:    ndx,
					xdir: 0,
					ydir: 0,
					next: NEXT_TOKEN_VALUE[j],
				})
			}
		}
	}

	//stage 2 - check all "M"'s for adjacent "A"'s

	var matches [][]xmastoken
	for _, x := range startlocs {
		if c := check(matrix, x); c != nil {
			matches = append(matches, c...)
		}
	}

	//stage 3 - follow each path for two more steps to check if full word is present

	var fullmatches [][]int
	for _, match := range matches {
		// lm := match[0]
		a := match[1]
		movedx := a.x + a.xdir
		movedy := a.y + a.ydir

		if movedx < 0 || movedy < 0 {
			continue
		}

		//this is ugly
		if len(matrix) > movedx {
			if len(matrix[movedx]) > movedy {
				if matrix[movedx][movedy] == a.next {
					//we matched an "S"
					// matchedS := xmastoken{
					// 	x:    movedx,
					// 	y:    movedy,
					// 	xdir: m.xdir,
					// 	ydir: m.ydir,
					// 	next: NEXT_TOKEN_VALUE[m.next],
					// }
					// var newnew []xmastoken
					// newnew = append(newnew, matches[idx]...)
					// newnew = append(newnew, matchedS)
					// fullmatches = append(fullmatches, newnew)
					fullmatches = append(fullmatches, []int{a.x, a.y})
				}
			}
		}

	}

	goodgoods := 0
	for idx, f := range fullmatches {
		for ndx, g := range fullmatches {
			if idx == ndx {
				continue
			}

			if slices.Equal[[]int](f, g) {
				goodgoods++
			}
			// a1 := f[1]
			// a2 := g[1]

			// if a1.x == a2.x && a1.y == a2.y {
			// 	goodgoods++
			// }
		}
	}

	goodgoods = goodgoods / 2 //each x is counted twice in loops above since they are not removed as they are matched

	log.Printf("%d words found\n", goodgoods)
	log.Printf("Execution time: %s", time.Since(st).String())
}

// checks all adjacent spots for values that match base.next and returns a slice for each
// possible path
func check(mtrx [][]rune, base xmastoken) [][]xmastoken {
	var matches [][]xmastoken
	//loop thru all surrounding spots
	for _, dir := range SURROUNDINGS {
		xd := dir[0]
		yd := dir[1]
		movedx := base.x + xd
		movedy := base.y + yd

		if movedx < 0 || movedy < 0 {
			continue
		}

		if len(mtrx) > movedx {
			if len(mtrx[movedx]) > movedy {
				if mtrx[movedx][movedy] == base.next {
					base.xdir = xd
					base.ydir = yd
					match := []xmastoken{
						base,
						{
							x:    movedx,
							y:    movedy,
							xdir: xd,
							ydir: yd,
							next: NEXT_TOKEN_VALUE[mtrx[movedx][movedy]],
						},
					}

					matches = append(matches, match)
				}
			}
		}
	}

	if len(matches) == 0 {
		return nil
	} else {
		return matches
	}
}

//Execution time: 41.1451ms
