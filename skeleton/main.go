package main

import (
	"aoc/util/aocutil"
	_ "embed"
	"fmt"
	"os"
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

	fmt.Printf("%d - %d\n", input.LineCount, input.CharCount)

	return solution
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	fmt.Printf("%d - %d\n", input.LineCount, input.CharCount)

	return solution
}
