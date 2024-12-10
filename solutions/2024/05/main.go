package main

import (
	"aoc/util/aocutil"
	"aoc/util/sliceutil"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"slices"
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
	rules, lists := getRulesAndLists()

	for _, l := range lists {
		steps := strings.Split(l, ",")

		correct := true
		for i := 0; i < len(steps)-1; i++ {
			found := false
			for _, r := range rules {
				if r == fmt.Sprintf("%s|%s", steps[i], steps[i+1]) {
					found = true
				}
			}

			if !found {
				correct = false
			}
		}

		if correct {
			nums := strings.Split(l, ",")
			mid, _ := strconv.Atoi(nums[(len(nums)-1)/2])
			solution += mid
		}
	}

	return solution
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)
	rules, lists := getRulesAndLists()

	processedRules := mapRules(rules)

	for _, l := range lists {
		steps := strings.Split(l, ",")

		correct := true
		for i := 0; i < len(steps)-1; i++ {
			found := false
			for _, r := range rules {
				if r == fmt.Sprintf("%s|%s", steps[i], steps[i+1]) {
					found = true
				}
			}

			if !found {
				correct = false
			}
		}

		if !correct {
			strNums := strings.Split(l, ",")

			var intNums []int
			for _, n := range strNums {
				converted, _ := strconv.Atoi(n)
				intNums = append(intNums, converted)
			}

			// check elements in order
			// if a preceding element is present in the processed rule map
			// for the current item, than swap them
			for j := 0; j < len(intNums); j++ {
				if j == 0 {
					continue
				}

				rulesForNum := processedRules[intNums[j]]
				for k := 0; k < j; k++ {
					// c, _ := sliceutil.IntContainsElement(rulesForNum, intNums[k])
					c := slices.Index(rulesForNum, intNums[k])
					if c != -1 {
						sliceutil.SwapSliceElements(intNums, j, k)
					}
				}
			}

			solution += intNums[(len(intNums)-1)/2]
		}
	}

	return solution
}

func getRulesAndLists() (rules []string, lists []string) {
	for _, line := range input.Lines {
		if line != "" {
			if line[2] == '|' {
				rules = append(rules, line)
			} else {
				lists = append(lists, line)
			}
		}
	}

	return rules, lists
}

func mapRules(rules []string) map[int][]int {
	processedRules := make(map[int][]int)
	for _, r := range rules {
		parts := strings.Split(r, "|")
		key, _ := strconv.Atoi(parts[0])
		val, _ := strconv.Atoi(parts[1])

		processedRules[key] = append(processedRules[key], val)
	}

	return processedRules
}
