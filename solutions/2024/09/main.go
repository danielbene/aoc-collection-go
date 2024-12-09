package main

import (
	"aoc/util/aocutil"
	_ "embed"
	"fmt"
	"os"
	"strconv"
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

	var diskLayout strings.Builder

	blockId := 0
	for _, line := range input.Lines {
		for idx, ch := range line {
			num, _ := strconv.Atoi(string(ch))
			if idx%2 == 0 {
				//block
				for i := 0; i < num; i++ {
					diskLayout.WriteString(strconv.Itoa(blockId))
				}

				blockId++
			} else {
				for i := 0; i < num; i++ {
					diskLayout.WriteString(".")
				}
			}
		}
	}

	fmt.Println(diskLayout.String())

	ordered := []rune(diskLayout.String())
	for i := len(ordered) - 1; i >= 0; i-- {
		if ordered[i] == '.' {
			continue
		}

		for j := 0; j < len(ordered); j++ {
			if j > i {
				break
			}

			if ordered[j] == '.' {
				ordered[j], ordered[i] = ordered[i], ordered[j]
			}
		}

	}

	for idx, ch := range ordered {
		if ch == '.' {
			return
		}

		num, _ := strconv.Atoi(string(ch))
		fmt.Printf("%d * %d\n", idx, num)

		solution += idx * num
	}

	//fmt.Println(diskLayout.String())
	//fmt.Println(string(ordered))

	return solution
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	fmt.Printf("%d - %d\n", input.LineCount, input.CharCount)

	return solution
}
