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

	// boooyah
	for y, line := range input.Lines {
		for x, ch := range line {
			if rune(ch) != '.' {
				for y2, line2 := range input.Lines {
					for x2, ch2 := range line2 {
						if ch2 == ch && (y != y2 || x != x2) {
							antiX := x - (x2 - x)
							antiY := y - (y2 - y)

							fmt.Printf("orig: %d %d secondary: %d %d anti: %d %d rune: %s\n", x, y, x2, y2, antiX, antiY, string(ch))
							if antiX >= 0 && antiY >= 0 &&
								antiX < input.CharCount && antiY < input.LineCount &&
								input.Lines[antiY][antiX] == '.' {
								//fmt.Printf("orig: %d %d secondary: %d %d anti: %d %d rune: %s target: %s\n", x, y, x2, y2, antiX, antiY, string(ch), string(input.Lines[antiY][antiX]))
								fmt.Println("ok")
								solution++
							}
						}
					}
				}
			}
		}
	}

	return solution
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	fmt.Printf("%d - %d\n", input.LineCount, input.CharCount)

	return solution
}
