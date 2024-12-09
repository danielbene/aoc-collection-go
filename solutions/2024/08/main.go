package main

import (
	"aoc/util/aocutil"
	_ "embed"
	"fmt"
	"os"
	"slices"
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

	// booooyah
	var coords []string
	for y, line := range input.Lines {
		for x, ch := range line {
			if ch != '.' {
				for y2, line2 := range input.Lines {
					for x2, ch2 := range line2 {
						if ch2 == ch && !(y == y2 && x == x2) {
							antiX := x - (x2 - x)
							antiY := y - (y2 - y)

							if antiX >= 0 && antiY >= 0 && antiX < input.CharCount && antiY < input.LineCount {
								c := fmt.Sprintf("%d,%d", antiX, antiY)
								if !slices.Contains(coords, c) {
									coords = append(coords, c)
								}
							}
						}
					}
				}
			}
		}
	}

	return len(coords)
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	// booooyah
	var coords []string
	for y, line := range input.Lines {
		for x, ch := range line {
			if ch != '.' {
				for y2, line2 := range input.Lines {
					for x2, ch2 := range line2 {
						if ch2 == ch && !(y == y2 && x == x2) {
							coords = append(coords, getAntinodes(x, y, x2, y2)...)
						}
					}
				}
			}
		}
	}

	return len(coords)
}

func getAntinodes(x int, y int, x2 int, y2 int) []string {
	var coords []string

	c := fmt.Sprintf("%d,%d", x, y)
	if !slices.Contains(coords, c) {
		coords = append(coords, c)
	}

	c = fmt.Sprintf("%d,%d", x2, y2)
	if !slices.Contains(coords, c) {
		coords = append(coords, c)
	}

	antiX := x - (x2 - x)
	antiY := y - (y2 - y)

	if antiX >= 0 && antiY >= 0 && antiX < input.CharCount && antiY < input.LineCount {
		c := fmt.Sprintf("%d,%d", antiX, antiY)
		if !slices.Contains(coords, c) {
			coords = append(coords, fmt.Sprintf("%d,%d", antiX, antiY))
		}
	}

	return coords
}
