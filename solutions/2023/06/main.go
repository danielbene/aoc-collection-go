package main

import (
	"aoc/util/aocutil"
	_ "embed"
	"fmt"
	"os"
	"regexp"
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
	solution = 1

	r := regexp.MustCompile(`\d+`)
	times := r.FindAllStringSubmatch(input.Lines[0], -1)
	distances := r.FindAllStringSubmatch(input.Lines[1], -1)

	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i][0])
		d, _ := strconv.Atoi(distances[i][0])

		farther := 0
		for j := t; j > 0; j-- {
			if j*(t-j) > d {
				farther++
			}
		}

		if farther > 0 {
			solution *= farther
		}
	}

	return solution
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	time := strings.ReplaceAll(strings.Replace(input.Lines[0], "Time:", "", 1), " ", "")
	distance := strings.ReplaceAll(strings.Replace(input.Lines[1], "Distance:", "", 1), " ", "")

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)

	for j := t; j > 0; j-- {
		if j*(t-j) > d {
			solution++
		}
	}

	return solution
}
