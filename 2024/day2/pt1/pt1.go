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
	for _, l := range lines {
		ld := strings.Fields(l) //cant Split() on whitespace, using Fields will split on whitespace
		last := 0
		increasing := false
		decreasing := false
		safe := true

		for idx, d := range ld {
			nmbr, _ := strconv.Atoi(d)
			if idx == 0 {
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
				break
			}

			if math.Abs(float64(last-nmbr)) > 3 {
				safe = false
				break
			}
			last = nmbr
		}

		if increasing && decreasing {
			safe = false
			increasing = false
			decreasing = false
		}

		if safe {
			safeCount++
		}
	}

	log.Printf("%d safe reports\n", safeCount)
	log.Printf("Execution time: %s", time.Since(st).String())
}

//Execution time: 510.9µs
