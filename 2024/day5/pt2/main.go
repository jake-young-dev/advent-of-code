package main

//Explanation
// This problem was solved by converting the rule data set to vertices and inverse them then calculate the adjacency matrix of our updates
// marking all points as connected with one edge. If the rule vertex has any edges in the adjacency matrix then the update is bad

import (
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	newline       = "\r\n"
	doublenewline = "\r\n\r\n"
	ruledivider   = "|"
)

func main() {
	st := time.Now()

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fbytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fstr := string(fbytes)

	filesplit := strings.Split(fstr, doublenewline)
	if len(filesplit) < 2 {
		panic("could not split file")
	}

	rulestr := filesplit[0]
	updatestr := filesplit[1]

	rules := strings.Split(rulestr, newline)
	type ruleobj struct {
		x int
		y int
	}

	//create ruleset vertices (inversed values)
	var ruleset []ruleobj
	for _, rr := range rules {
		rp := strings.Split(rr, ruledivider)
		x, _ := strconv.Atoi(rp[1])
		y, _ := strconv.Atoi(rp[0])
		ruleset = append(ruleset, ruleobj{x: x, y: y})
	}

	//create adjacency matrix for each update
	updates := strings.Split(updatestr, newline)
	finalCount := 0
	badUpdates := make([][]int, 0)
	for _, up := range updates {
		adjmatrix := make(map[int]map[int]int, 0)
		dstr := strings.Split(up, ",")
		d, err := convertToInt(dstr)
		if err != nil {
			panic(err)
		}
		for idx, a := range d {
			for ndx := idx; ndx < len(d); ndx++ {
				if idx == ndx {
					continue
				}
				if _, ok := adjmatrix[a]; !ok {
					adjmatrix[a] = make(map[int]int, 0)
				}
				if _, ok := adjmatrix[a][d[ndx]]; !ok {
					adjmatrix[a][d[ndx]] = 1
				}
			}
		}

		//if vertex from ruleset has edges in update matrix the update needs to be discarded
		for _, r := range ruleset {
			val, ok := adjmatrix[r.x][r.y]
			if ok {
				if val == 1 {
					badUpdates = append(badUpdates, d)
					break
				}
			}
		}
	}

	slices.Reverse(ruleset)
	for _, bd := range badUpdates {
		for _, r := range ruleset {
			slices.SortStableFunc(bd, func(a, b int) int {
				if a == b {
					return 0
				}
				if r.x == b {
					if r.y == a {
						return -1
					}
				}
				return 1
			})
		}

		finalCount += bd[len(bd)/2]
	}

	log.Printf("%d found\n", finalCount)
	log.Printf("Execution time: %s", time.Since(st).String())
}

func convertToInt(old []string) ([]int, error) {
	new := make([]int, len(old))
	for i, s := range old {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		new[i] = num
	}

	return new, nil
}

//Execution time: 22.5897ms
