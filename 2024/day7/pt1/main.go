package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type tree struct {
	value int
	left  *tree
	right *tree
}

func main() {
	st := time.Now()
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fbytes, err := io.ReadAll(file)
	fstr := string(fbytes)

	lines := strings.Split(fstr, "\r\n")

	count := 0
	for _, l := range lines {
		splLine := strings.Split(l, ": ")
		if len(splLine) < 2 {
			panic(errors.New("incomplete piece"))
		}
		//cache total
		total := splLine[0]
		piecestr := splLine[1]

		//split can be unpredictable with Split()
		pieces := strings.Fields(piecestr)

		firstitem, _ := strconv.Atoi(pieces[0])

		//create root node for current tree
		root := &tree{
			value: firstitem,
			left:  nil,
			right: nil,
		}

		//recursively create binary tree
		splitNode(root, 1, pieces)

		//grab all leaves
		var solutions []int
		search(root, &solutions)

		//if leaf is total we found a possible solution
		tt, _ := strconv.Atoi(total)
		for _, s := range solutions {
			if s == tt {
				count += tt
				break
			}
		}
	}

	log.Printf("%d total calibrations found\n", count)
	log.Printf("Execution time: %s", time.Since(st).String())
}

// depth-first-style search
func search(node *tree, solution *[]int) {
	if node.left == nil && node.right == nil {
		*solution = append(*solution, node.value)
	}
	if node.left != nil {
		search(node.left, solution)
	}
	if node.right != nil {
		search(node.right, solution)
	}
}

// creates children for node with data from pieces indexed at i
func splitNode(node *tree, i int, pieces []string) {
	if i >= len(pieces) {
		return
	}
	num, _ := strconv.Atoi(pieces[i])
	nl := &tree{
		value: node.value + num,
		left:  nil,
		right: nil,
	}
	node.left = nl
	nr := &tree{
		value: node.value * num,
		left:  nil,
		right: nil,
	}
	node.right = nr
	splitNode(nr, i+1, pieces)
	splitNode(nl, i+1, pieces)
}

//Execution time: 31.6272ms
