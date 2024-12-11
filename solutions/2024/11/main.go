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

func blink_o_matic(count int) int {
	var stones []uint64
	for _, line := range input.Lines {
		tmp := strings.Split(line, " ")
		for _, num := range tmp {
			n, _ := strconv.Atoi(num)
			stones = append(stones, uint64(n))
		}
	}

	for i := 0; i < count; i++ {
		fmt.Printf("loop: %d\n", i)
		blink(&stones)
	}

	return len(stones)
}

func blink(stones *[]uint64) {
	for idx, stone := range *stones {
		strStone := strconv.Itoa(int((*stones)[idx]))
		digitCnt := len(strStone)
		if stone == 0 {
			(*stones)[idx] = 1
		} else if digitCnt%2 == 0 {
			a, _ := strconv.Atoi(strStone[:digitCnt/2])
			b, _ := strconv.Atoi(strStone[digitCnt/2:])

			(*stones)[idx] = uint64(a)
			*stones = append(*stones, uint64(b)) // putting in the back to avoid idx duckup
		} else {
			(*stones)[idx] = (*stones)[idx] * 2024
		}
	}
}
