package main

import (
	"aoc/util"
	"aoc/util/sliceutil"
	_ "embed"
	"fmt"
	"math"
	"os"
	"regexp"
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
		util.Solve(p1, time.Since(start), "Part1", os.Args[2])
	} else {
		fmt.Println("running part2")
		p2 := Part2(inputFile)
		util.Solve(p2, time.Since(start), "Part2", os.Args[2])
	}
}

// -----------------------------------------------------------

var input util.ProcessedInput

func Part1(puzzleInput string) (solution int) {
	util.ProcessInput(puzzleInput, &input)

	for _, line := range input.Lines {
		parts := strings.Split(line, ": ")
		sides := strings.Split(parts[1], " | ")

		var winningNums []string
		var myNums []string

		r := regexp.MustCompile(`\d+`)
		for _, val := range r.FindAllStringSubmatch(sides[0], -1) {
			winningNums = append(winningNums, val...)
		}

		for _, val := range r.FindAllStringSubmatch(sides[1], -1) {
			myNums = append(myNums, val...)
		}

		commonElementsCnt := sliceutil.CountCommonSliceElements(winningNums, myNums)
		if commonElementsCnt > 0 {
			solution += int(math.Pow(2, float64((commonElementsCnt - 1))))
		}
	}

	return solution
}

func Part2(puzzleInput string) (solution int) {
	util.ProcessInput(puzzleInput, &input)
	copies := make(map[int]int) // cardId : numOfCopies

	for currentLineId, line := range input.Lines {
		parts := strings.Split(line, ": ")
		sides := strings.Split(parts[1], " | ")

		var winningNums []string
		var myNums []string

		r := regexp.MustCompile(`\d+`)
		for _, val := range r.FindAllStringSubmatch(sides[0], -1) {
			winningNums = append(winningNums, val...)
		}

		for _, val := range r.FindAllStringSubmatch(sides[1], -1) {
			myNums = append(myNums, val...)
		}

		commonElementsCnt := sliceutil.CountCommonSliceElements(winningNums, myNums)
		for i := currentLineId; i < currentLineId+commonElementsCnt; i++ {
			copies[i+1] += 1 + copies[currentLineId]
		}
	}

	for _, v := range copies {
		solution += v
	}

	solution += input.LineCount // add the original cards

	return solution
}
