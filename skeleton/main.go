package main

import (
	"aoc/util"
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
}

func main() {
	start := time.Now()
	if os.Args[1] == "1" {
		fmt.Println("running part1")
		p1 := Part1(input)
		util.Solve(p1, time.Since(start), "Part1", os.Args[2])
	} else {
		fmt.Println("running part2")
		p2 := Part2(input)
		util.Solve(p2, time.Since(start), "Part2", os.Args[2])
	}
}

// -----------------------------------------------------------

var (
	lines     []string
	lineCount = 0
	charCount = 0
)

func Part1(puzzleInput string) int {
	processInput(puzzleInput)

	fmt.Printf("%d - %d\n", lineCount, charCount)

	return 0
}

func Part2(puzzleInput string) int {
	processInput(puzzleInput)

	fmt.Printf("%d - %d\n", lineCount, charCount)

	return 0
}

func processInput(puzzleInput string) {
	if len(lines) != 0 {
		fmt.Println("Input already processed.")
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	lineCount = len(lines)
	charCount = len(lines[0])
}
