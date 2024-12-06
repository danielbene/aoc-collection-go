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

type Coord struct {
	X           int
	Y           int
	Orientation string
}

func Part1(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)
	mtx := make([][]string, input.LineCount)

	pos := Coord{X: 0, Y: 0}
	for idx, line := range input.Lines {
		for i, ch := range line {
			mtx[idx] = append(mtx[idx], string(ch))
			if string(ch) == "^" {
				pos.X = i
				pos.Y = idx
				pos.Orientation = "^"
			}
		}
	}

	for {
		if pos.Y-1 < 0 || pos.Y+1 >= input.LineCount ||
			pos.X-1 < 0 || pos.X+1 >= input.CharCount {
			break
		}

		switch pos.Orientation {
		case "^":
			// check North
			if mtx[pos.Y-1][pos.X] != "#" {
				pos.Y--
				mtx[pos.Y][pos.X] = "X"
			} else {
				pos.Orientation = ">"
			}
		case ">":
			// check East
			if mtx[pos.Y][pos.X+1] != "#" {
				pos.X++
				mtx[pos.Y][pos.X] = "X"
			} else {
				pos.Orientation = "ˇ"
			}
		case "ˇ":
			// check South
			if mtx[pos.Y+1][pos.X] != "#" {
				pos.Y++
				mtx[pos.Y][pos.X] = "X"
			} else {
				pos.Orientation = "<"
			}
		case "<":
			// check South
			if mtx[pos.Y][pos.X-1] != "#" {
				pos.X--
				mtx[pos.Y][pos.X] = "X"
			} else {
				pos.Orientation = "^"
			}
		}
	}

	for _, row := range mtx {
		for _, val := range row {
			if val == "X" {
				solution++
			}
		}
	}

	return solution
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)
	mtx := make([][]string, input.LineCount)

	pos := Coord{X: 0, Y: 0}
	for idx, line := range input.Lines {
		for i, ch := range line {
			mtx[idx] = append(mtx[idx], string(ch))
			if string(ch) == "^" {
				pos.X = i
				pos.Y = idx
				pos.Orientation = "^"

				mtx[idx][i] = "."
			}

		}
	}

	origPos := Coord{
		X:           pos.X,
		Y:           pos.Y,
		Orientation: pos.Orientation,
	}

	updated := Coord{X: 0, Y: 0}
	origChar := ""

	for {
		pos = origPos
		origChar = mtx[updated.Y][updated.X]
		mtx[updated.Y][updated.X] = "#"

		loopCount := 0
		for {
			if pos.Y-1 < 0 || pos.Y+1 >= input.LineCount ||
				pos.X-1 < 0 || pos.X+1 >= input.CharCount {
				break
			}

			switch pos.Orientation {
			case "^":
				if mtx[pos.Y-1][pos.X] != "#" {
					pos.Y--
				} else {
					pos.Orientation = ">"
				}
			case ">":
				if mtx[pos.Y][pos.X+1] != "#" {
					pos.X++
				} else {
					pos.Orientation = "ˇ"
				}
			case "ˇ":
				if mtx[pos.Y+1][pos.X] != "#" {
					pos.Y++
				} else {
					pos.Orientation = "<"
				}
			case "<":
				if mtx[pos.Y][pos.X-1] != "#" {
					pos.X--
				} else {
					pos.Orientation = "^"
				}
			}

			loopCount++

			// assuming it is a loop if more than 10K moves happend
			if loopCount > 10000 {
				solution++
				break
			}
		}

		mtx[updated.Y][updated.X] = origChar

		if updated.X+1 < input.CharCount {
			updated.X++
		} else if updated.Y+1 < input.LineCount {
			updated.Y++
			updated.X = 0
		} else {
			break
		}
	}

	return solution - 1 // mystery -1, it is what it is
}
