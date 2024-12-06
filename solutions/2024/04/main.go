package main

import (
	"aoc/util/aocutil"
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
		aocutil.Solve(p1, time.Since(start), "Part1", os.Args[2])
	} else {
		fmt.Println("running part2")
		p2 := Part2(input)
		aocutil.Solve(p2, time.Since(start), "Part2", os.Args[2])
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
	count := 0

	for currentLineId, line := range lines {
		horizontal(&count, &line)
		for idx, ch := range line {
			if ch == 'X' || ch == 'S' {
				diagonal(&count, currentLineId, idx, ch)
			}
		}
	}

	return count
}

func Part2(puzzleInput string) int {
	processInput(puzzleInput)
	count := 0

	for currentLineId, line := range lines {
		if currentLineId >= 1 && currentLineId < lineCount-1 {
			for idx, ch := range line {
				if idx > 0 && idx < charCount-1 {
					if ch == 'A' {
						if checkWings(currentLineId, idx) {
							count++
						}
					}
				}
			}
		}
	}

	return count
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

func checkWings(currentLineId int, currentCharId int) bool {
	leftUp := lines[currentLineId-1][currentCharId-1]
	leftDown := lines[currentLineId+1][currentCharId-1]
	rightUp := lines[currentLineId-1][currentCharId+1]
	rightDown := lines[currentLineId+1][currentCharId+1]

	if (leftUp == 'M' && rightDown == 'S') || (leftUp == 'S' && rightDown == 'M') {
		if (leftDown == 'M' && rightUp == 'S') || (leftDown == 'S' && rightUp == 'M') {
			return true
		}
	}

	return false
}

func diagonal(count *int, currentLineId int, currentCharId int, firstRune rune) {
	var target string
	if firstRune == 'X' {
		target = "MAS"
	} else {
		target = "AMX"
	}

	// NOTE: rightUp and leftUp does not needed because we checking both ways
	verti, leftDown, rightDown := 1, 1, 1
	for i, req := range target {
		// vertical
		if verti == 1 {
			targetLineId := currentLineId + i + 1
			if targetLineId >= lineCount || rune(lines[targetLineId][currentCharId]) != req {
				verti = 0
			}
		}

		// left-down
		row := currentLineId + i + 1
		col := currentCharId - i - 1
		if leftDown == 1 {
			if !(row < lineCount && col >= 0) || rune(lines[row][col]) != req {
				leftDown = 0
			}
		}

		// right-down
		row = currentLineId + i + 1
		col = currentCharId + i + 1
		if rightDown == 1 {
			if !(row < lineCount && col < charCount) || rune(lines[row][col]) != req {
				rightDown = 0
			}
		}
	}

	*count += verti + leftDown + rightDown
}

func horizontal(count *int, line *string) {
	*count += strings.Count(*line, "XMAS")
	*count += strings.Count(*line, "SAMX")
}
