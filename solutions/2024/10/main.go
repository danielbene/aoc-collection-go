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
var counter int
var visited []string

func Part1(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)
	matrix := getMatrix()

	counter = 0
	for y, row := range matrix {
		for x, num := range row {
			if num == 0 {
				visited = []string{}
				checkNext(x, y, 0, matrix, false)
			}
		}
	}

	return counter
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)
	matrix := getMatrix()

	counter = 0
	for y, row := range matrix {
		for x, num := range row {
			if num == 0 {
				checkNext(x, y, 0, matrix, true)
			}
		}
	}

	return counter
}

func getMatrix() [][]int {
	matrix := make([][]int, input.LineCount)
	for y, line := range input.Lines {
		matrix[y] = make([]int, input.CharCount)
		for x, ch := range line {
			num := int(ch - '0')
			matrix[y][x] = num
		}
	}

	return matrix
}

// param - current x,y, and field value
func checkNext(x int, y int, val int, matrix [][]int, countDistinct bool) {
	defer func() {
		if r := recover(); r != nil {
			return // lazy handling of index out of bounds
		}
	}()

	if matrix[y][x] == val {
		if val == 9 {
			coord := fmt.Sprintf("%d,%d", x, y)
			if !slices.Contains(visited, coord) || countDistinct {
				visited = append(visited, coord)
				counter++
			}

			return
		}

		// no diagonals
		checkNext(x, y-1, val+1, matrix, countDistinct)
		checkNext(x-1, y, val+1, matrix, countDistinct)
		checkNext(x, y+1, val+1, matrix, countDistinct)
		checkNext(x+1, y, val+1, matrix, countDistinct)
	}
}
