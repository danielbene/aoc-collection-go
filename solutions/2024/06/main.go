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
	mtx, pos := getMatrixAndPos()

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
	mtx, currGuardPos := getMatrixAndPos()

	startGuardPos := Coord{
		X:           currGuardPos.X,
		Y:           currGuardPos.Y,
		Orientation: currGuardPos.Orientation,
	}

	testBlockCoord := Coord{X: 0, Y: 0}
	origChar := ""

	// basically testing each coord with a temporary block and trying to find loops
	for {
		currGuardPos = startGuardPos

		origChar = mtx[testBlockCoord.Y][testBlockCoord.X]
		mtx[testBlockCoord.Y][testBlockCoord.X] = "#"

		loopCount := 0
		for {
			if currGuardPos.Y-1 < 0 || currGuardPos.Y+1 >= input.LineCount ||
				currGuardPos.X-1 < 0 || currGuardPos.X+1 >= input.CharCount {
				break
			}

			switch currGuardPos.Orientation {
			case "^":
				if mtx[currGuardPos.Y-1][currGuardPos.X] != "#" {
					currGuardPos.Y--
				} else {
					currGuardPos.Orientation = ">"
				}
			case ">":
				if mtx[currGuardPos.Y][currGuardPos.X+1] != "#" {
					currGuardPos.X++
				} else {
					currGuardPos.Orientation = "ˇ"
				}
			case "ˇ":
				if mtx[currGuardPos.Y+1][currGuardPos.X] != "#" {
					currGuardPos.Y++
				} else {
					currGuardPos.Orientation = "<"
				}
			case "<":
				if mtx[currGuardPos.Y][currGuardPos.X-1] != "#" {
					currGuardPos.X--
				} else {
					currGuardPos.Orientation = "^"
				}
			}

			loopCount++

			// assuming it is a loop if more than X moves happend
			if loopCount > 100000 {
				solution++
				break
			}
		}

		mtx[testBlockCoord.Y][testBlockCoord.X] = origChar

		if testBlockCoord.X+1 < input.CharCount {
			testBlockCoord.X++

			// cannot block starting position
			if testBlockCoord.Y == startGuardPos.Y && testBlockCoord.X == startGuardPos.X {
				testBlockCoord.X++
			}
		} else if testBlockCoord.Y+1 < input.LineCount {
			testBlockCoord.Y++
			testBlockCoord.X = 0
		} else {
			break
		}
	}

	// secondary input solution should be 1789
	return solution
}

func getMatrixAndPos() ([][]string, Coord) {
	mtx := make([][]string, input.LineCount)

	startGuardPos := Coord{X: 0, Y: 0}
	for idx, line := range input.Lines {
		for i, ch := range line {
			mtx[idx] = append(mtx[idx], string(ch))
			if string(ch) == "^" {
				startGuardPos.X = i
				startGuardPos.Y = idx
				startGuardPos.Orientation = "^"

				mtx[idx][i] = "."
			}
		}
	}

	return mtx, startGuardPos
}
