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

// TODO: int return usage in skeleton

const (
	xmas = "XMAS"
	samx = "SAMX"
)

func Part1(puzzleInput string) int {
	var (
		lines     []string
		count     = 0
		lineCount = 0
		charCount = 0
	)

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	lineCount = len(lines)
	charCount = len(lines[0])

	diagCnt := 0
	for currentLineId, line := range lines {
		horizontal(&count, &line)
		for idx, ch := range line {
			if ch == 'X' || ch == 'S' {
				vertical(&count, currentLineId, idx, ch, lines, lineCount)
				diagonal(&diagCnt, currentLineId, idx, ch, lines, lineCount, charCount)
			}
		}
	}

	count += diagCnt / 2 // counted to and from so we have to halve the value

	return count
}

func diagonal(count *int, currentLineId int, currentCharId int, firstRune rune, lines []string, lineCount int, charCount int) {
	var target string
	if firstRune == 'X' {
		target = "MAS"
	} else {
		target = "AMX"
	}

	leftUp, leftDown, rightUp, rightDown := true, true, true, true
	for i, req := range target {
		// left-up
		row := currentLineId - i - 1
		col := currentCharId - i - 1
		if leftUp {
			if !(row >= 0 && col >= 0) || rune(lines[row][col]) != req {
				leftUp = false
			}
		}

		// left-down
		row = currentLineId + i + 1
		col = currentCharId - i - 1
		if leftDown {
			if !(row < lineCount && col >= 0) || rune(lines[row][col]) != req {
				leftDown = false
			}
		}

		// right-up
		row = currentLineId - i - 1
		col = currentCharId + i + 1
		if rightUp {
			if !(row >= 0 && col < charCount) || rune(lines[row][col]) != req {
				rightUp = false
			}
		}

		// right-down
		row = currentLineId + i + 1
		col = currentCharId + i + 1
		if rightDown {
			if !(row < lineCount && col < charCount) || rune(lines[row][col]) != req {
				rightDown = false
			}
		}
	}

	if leftUp {
		*count++
	}

	if leftDown {
		*count++
	}

	if rightUp {
		*count++
	}

	if rightDown {
		*count++
	}
}

func horizontal(count *int, line *string) {
	*count += strings.Count(*line, xmas) // non overlapping count?
	*count += strings.Count(*line, samx) // non overlapping count?
}

func vertical(count *int, currentLineId int, currentCharId int, firstRune rune, lines []string, lineCount int) {
	var target string
	if firstRune == 'X' {
		target = "MAS"
	} else {
		target = "AMX"
	}

	for i, req := range target {
		if currentLineId+i+1 >= lineCount {
			return
		}

		if rune(lines[currentLineId+i+1][currentCharId]) != req {
			return
		}
	}

	*count++
}

func Part2(puzzleInput string) int {
	var (
		lines     []string
		count     = 0
		lineCount = 0
		charCount = 0
	)

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	lineCount = len(lines)
	charCount = len(lines[0])

	for currentLineId, line := range lines {
		if currentLineId >= 1 && currentLineId < lineCount-1 {
			for idx, ch := range line {
				if idx > 0 && idx < charCount-1 {
					if ch == 'A' {
						if checkWings(currentLineId, idx, lines) {
							count++
						}
					}
				}
			}
		}
	}

	return count
}

func checkWings(currentLineId int, currentCharId int, lines []string) bool {
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
