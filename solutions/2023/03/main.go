package main

import (
	"aoc/util"
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
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

func Part1(puzzleInput string) string {
	var (
		lines []string
		nums  []int
		sum   int
	)

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lineCnt := len(lines)
	for currentLineId, line := range lines {
		xCnt := len(line)

		continueIdx := 0 // handling dup nums in the same line - needed because index usage
		r := regexp.MustCompile(`\d+`)
		for _, val := range r.FindAllStringSubmatch(line, -1) {
			start := strings.Index(line[continueIdx:], val[0]) + continueIdx
			end := start + len(val[0])
			continueIdx = end

			dec, _ := strconv.Atoi(val[0])

			// before
			if start != 0 {
				if checkSpec(line[start-1]) {
					nums = append(nums, dec)
					continue
				}
			}

			// after
			if end < xCnt {
				if checkSpec(line[end]) {
					nums = append(nums, dec)
					continue
				}
			}

			// above
			if currentLineId > 0 {
				ok := false // avoid double addition
				for i := start - 1; i <= end; i++ {
					if i < 0 || i >= xCnt {
						continue
					}

					if checkSpec(lines[currentLineId-1][i]) {
						nums = append(nums, dec)
						ok = true
						break
					}
				}

				if ok {
					continue
				}
			}

			// below
			if currentLineId < lineCnt-1 {
				ok := false // avoid double addition
				for i := start - 1; i <= end; i++ {
					if i < 0 || i >= xCnt {
						continue
					}

					if checkSpec(lines[currentLineId+1][i]) {
						nums = append(nums, dec)
						ok = true
						break
					}
				}

				if ok {
					continue
				}
			}
		}
	}

	for _, num := range nums {
		sum += num
	}

	return strconv.Itoa(sum)
}

func Part2(puzzleInput string) string {
	var (
		lines []string
		sum   int
	)

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// lineCnt := len(lines)
	for currentLineId, line := range lines {
		xCnt := len(line)

		continueIdx := 0 // handling dup nums in the same line - needed because index usage
		r := regexp.MustCompile(`\*`)
		for _, val := range r.FindAllStringSubmatch(line, -1) {
			var gear []int

			start := strings.Index(line[continueIdx:], val[0]) + continueIdx
			end := start + 1
			continueIdx = end

			// before
			if start != 0 {
				if checkDigit(line[start-1]) {
					numStr := strings.Replace(string(line[start-3:start]), ".", "", -1)
					num, _ := strconv.Atoi(numStr)
					gear = append(gear, num)
				}
			}

			// after
			if end < xCnt {
				if checkDigit(line[end]) {
					numStr := strings.Replace(string(line[end:end+3]), ".", "", -1)
					num, _ := strconv.Atoi(numStr)
					gear = append(gear, num)
				}
			}

			// above
			if currentLineId > 0 {
				// for i := start - 1; i <= end; i++ {
				/*for i := start - 3; i <= end+2; i++ {
					if i < 0 || i >= xCnt {
						continue
					}

					fmt.Print(string(lines[currentLineId-1][i]))
				}
				fmt.Println()*/

				above := string(lines[currentLineId-1][start-3 : end+3])

				// if one of the middles is a digit than we good
				if checkDigit(above[2]) || checkDigit(above[3]) || checkDigit(above[4]) {
					fmt.Println(above)
					r := regexp.MustCompile(`\d+`)
					for _, val := range r.FindAllStringSubmatch(above, -1) {
						fmt.Print(val)
						fmt.Print(", ")
					}
					fmt.Println()
				}
			}

			// below
			/*if currentLineId < lineCnt-1 {
				// for i := start - 1; i <= end; i++ {
				for i := start - 3; i <= end+2; i++ {
					if i < 0 || i >= xCnt {
						continue
					}

					fmt.Print(string(lines[currentLineId+1][i]))
				}
				fmt.Println()
			}*/
		}
	}

	return strconv.Itoa(sum)
}

func checkSpec(ch byte) bool {
	return string(ch) != "." && !checkDigit(ch)
}

func checkDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

func checkAsterisk(ch byte) bool {
	return string(ch) == "*"
}
