package main

import (
	"aoc/util/aocutil"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat/combin"
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
	return combine(2)
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)
	return combine(3)
}

func combine(opCount int) (solution int) {
	for _, line := range input.Lines {
		parts := strings.Split(line, ": ")
		tValue, _ := strconv.Atoi(parts[0])
		operands := strings.Split(string(parts[1]), " ")

		setup := []int{}
		for i := 0; i < len(operands)-1; i++ {
			setup = append(setup, opCount)
		}

		comb := combin.Cartesian(setup)
		for _, c := range comb {
			var row strings.Builder
			for i, op := range operands {
				row.WriteString(op)

				if i < len(c) {
					if c[i] == 0 {
						row.WriteString("+")
					} else if c[i] == 1 {
						row.WriteString("*")
					} else {
						row.WriteString("@") // using @ instead of ||
					}
				}
			}

			if calculateRow(row.String()) == tValue {
				solution += tValue
				break
			}
		}
	}

	return solution
}

// always left-to-right
func calculateRow(inp string) int {
	var operators []rune
	numStrs := strings.FieldsFunc(inp, func(r rune) bool {
		if r == '+' || r == '*' || r == '@' {
			operators = append(operators, r)
			return true
		}

		return false
	})

	sum, _ := strconv.Atoi(numStrs[0])
	for idx, op := range operators {
		if idx+1 >= len(numStrs) {
			break
		}

		a, _ := strconv.Atoi(numStrs[idx+1])

		if op == '+' {
			sum += a
		} else if op == '*' {
			sum *= a
		} else if op == '@' {
			sum, _ = strconv.Atoi(fmt.Sprintf("%d%s", sum, numStrs[idx+1]))
		}
	}

	return sum
}
