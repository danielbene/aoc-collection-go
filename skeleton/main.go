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

func Part1(puzzleInput string) int {
	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return 0
}

func Part2(puzzleInput string) int {
	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return 0
}
