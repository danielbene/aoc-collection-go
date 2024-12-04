package main

import (
	"aoc/util/aocutil"
	"aoc/util/sliceutil"
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
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

func Part1(puzzleInput string) string {
	safeLevelCnt := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if checkLevels(parts) {
			safeLevelCnt++
		}
	}

	return strconv.Itoa(safeLevelCnt)
}

func Part2(puzzleInput string) string {
	safeLevelCnt := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		isLevelSafe := false
		parts := strings.Split(scanner.Text(), " ")

		isLevelSafe = checkLevels(parts)

		// a bit wasteful approach for toleration checks
		if !isLevelSafe {
			for i := 0; i < len(parts); i++ {
				tolerated := make([]string, len(parts))

				copy(tolerated, parts)
				tolerated = sliceutil.RemoveSliceElement(tolerated, i)

				isLevelSafe = checkLevels(tolerated)
				if isLevelSafe {
					break
				}
			}
		}

		if isLevelSafe {
			safeLevelCnt++
		}
	}

	return strconv.Itoa(safeLevelCnt)
}

func checkLevels(parts []string) bool {
	isCorrectDiff := true
	isDescending := true
	isAscending := true

	for i := 0; i < len(parts)-1; i++ {
		a, _ := strconv.Atoi(parts[i])
		b, _ := strconv.Atoi(parts[i+1])
		diff := a - b

		if isDescending && a < b {
			isDescending = false
		}

		if isAscending && a > b {
			isAscending = false
		}

		if !isAscending && !isDescending {
			break
		}

		if !((-3 <= diff && diff <= -1) || (3 >= diff && diff >= 1)) {
			isCorrectDiff = false
			break
		}
	}

	return isCorrectDiff && (isAscending || isDescending)
}
