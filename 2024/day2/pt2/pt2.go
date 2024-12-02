package main

import (
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	st := time.Now()
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
	lines := strings.Split(dataStr, "\r\n") //windows newline
	safeCount := 0
	var badlines []string //tracking lines that are "unsafe" for the "Problem Dampener" check
	for _, l := range lines {
		if safe := validate(l, -1); safe { //validate removing no items
			safeCount++
		} else {
			badlines = append(badlines, l)
		}
	}

	for _, l := range badlines {
		ld := strings.Fields(l)
		for x := range len(ld) { //validate after removing each index iteratively
			safe := validate(l, x)
			if safe {
				safeCount++
				break
			}
		}
	}

	log.Printf("%d safe reports\n", safeCount)
	log.Printf("Execution time: %s", time.Since(st).String())
}

func validate(l string, rm int) bool {
	ld := strings.Fields(l) //cant Split() on whitespace, using Fields will split on whitespace
	last := -1
	increasing := false
	decreasing := false
	safe := true

	for idx, d := range ld {
		nmbr, _ := strconv.Atoi(d)
		if idx == rm {
			continue
		}
		if idx == 0 {
			last = nmbr
			continue
		}
		if last == -1 {
			last = nmbr
			continue
		}

		if last < nmbr {
			increasing = true
		} else {
			decreasing = true
		}

		if last == nmbr {
			safe = false
		}

		if math.Abs(float64(last-nmbr)) > 3 {
			safe = false
		}
		last = nmbr
	}

	if increasing && decreasing {
		return false
	}

	if safe {
		return true
	}
	return false
}

//Execution time: 1.9946ms
