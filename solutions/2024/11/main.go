package main

import (
	"aoc/util/aocutil"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputFile string

func init() {
	inputFile = strings.TrimRight(inputFile, "\n")
}

func main() {
	start := time.Now()
	if os.Args[1] == "1" {
		fmt.Println("running part1")
		p1 := Part1(inputFile)
		aocutil.Solve(p1, time.Since(start), "Part1", os.Args[2])
	} else {
		fmt.Println("running part2")
		p2 := Part2(inputFile)
		aocutil.Solve(p2, time.Since(start), "Part2", os.Args[2])
	}
}

// -----------------------------------------------------------

var input aocutil.ProcessedInput

func Part1(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)
	return blink_o_matic(25)
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)
	return blink_o_matic(75)
}

func blink_o_matic(count int) (solution int) {
	var stones []uint64
	for _, line := range input.Lines {
		tmp := strings.Split(line, " ")
		for _, num := range tmp {
			n, _ := strconv.Atoi(num)
			stones = append(stones, uint64(n))
		}
	}

	// key: stone value, value: counter
	// instead of keeping track of individual stones we manage a ledger with stone counters
	ledger := make(map[uint64]int)
	for _, stone := range stones {
		ledger[stone]++
	}

	// doin da blinkers
	for i := 0; i < count; i++ {
		blink(&ledger)
	}

	for _, v := range ledger {
		solution += v
	}

	return solution
}

func blink(ledger *map[uint64]int) {
	keys := getStones(ledger)         // getting stones with at least one counter
	newStones := make(map[uint64]int) // tmp map for new stones to avoid inloop modification of the ledger

	for _, stone := range keys {
		strStone := strconv.Itoa(int(stone)) // probably not optimal but convenient
		digitCnt := len(strStone)

		cnt := (*ledger)[uint64(stone)]
		(*ledger)[uint64(stone)] = 0

		if stone == 0 {
			newStones[1] += cnt
		} else if digitCnt%2 == 0 {
			a, _ := strconv.Atoi(strStone[:digitCnt/2])
			b, _ := strconv.Atoi(strStone[digitCnt/2:])

			newStones[uint64(a)] += cnt
			newStones[uint64(b)] += cnt
		} else {
			newStones[uint64(stone*2024)] += cnt
		}
	}

	// update ledger with new stones
	for k, v := range newStones {
		(*ledger)[uint64(k)] += v
	}
}

func getStones(m *map[uint64]int) (keys []uint64) {
	for k, v := range *m {
		if v > 0 {
			keys = append(keys, k)
		}
	}

	return keys
}
