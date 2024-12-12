package main

import (
	"aoc/util/aocutil"
	"aoc/util/matrixutil"
	"aoc/util/matrixutil/directions"
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

var alreadyVisited [][]bool
var fenceCnt = 0
var fieldCnt = 0

func Part1(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	mtx := matrixutil.Init[string](input.Lines)
	mtx.Print()

	alreadyVisited = make([][]bool, mtx.RowCount)
	for i := range alreadyVisited {
		alreadyVisited[i] = make([]bool, mtx.ColCount)
	}

	for y, row := range mtx.Matrix {
		for x, ch := range row {
			if alreadyVisited[y][x] {
				continue
			}

			fieldCnt = 0
			fenceCnt = 0
			recursiveWalker(x, y, ch, &mtx)

			val := fieldCnt * fenceCnt
			fmt.Println(val)
			fmt.Println(fieldCnt)
			fmt.Println(fenceCnt)

			mtx.Print()

			for _, row := range alreadyVisited {
				fmt.Println(row)
			}

			return
		}
	}

	return solution
}

func recursiveWalker(x, y int, currVal string, mtx *matrixutil.Matrix[string]) {
	if !alreadyVisited[y][x] {
		if !mtx.Move(x, y) {
			panic("Achtung!")
		}

		fmt.Printf("x: %d, y: %d\n", x, y)
		fmt.Println(mtx.CurrentPosition)

		alreadyVisited[y][x] = true

		for _, dir := range directions.GetDirections() {
			nextX := x + dir.X
			nextY := y + dir.Y

			// FIXME: this does not check already visited!
			if val, success := mtx.GetValueDirection(dir); !success {
				fenceCnt++
			} else if val != currVal {
				fenceCnt++
			} else if val == currVal {
				fieldCnt++
				if mtx.IsInsideBoundaries(nextX, nextY) && !alreadyVisited[nextY][nextX] {
					recursiveWalker(nextX, nextY, currVal, mtx)
				}
			}
		}
	}
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	fmt.Printf("%d - %d\n", input.LineCount, input.CharCount)

	return solution
}
