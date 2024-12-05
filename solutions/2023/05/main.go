package main

import (
	"aoc/util/aocutil"
	_ "embed"
	"fmt"
	"math"
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

type Pair struct {
	Left  string
	Right *[][]int
}

var (
	seeds                 []int
	seedToSoil            [][]int
	soilToFertilizer      [][]int
	fertilizerToWater     [][]int
	waterToLight          [][]int
	lightToTemperature    [][]int
	temperatureToHumidity [][]int
	humidityToLocation    [][]int

	populationDone bool = false
)

var sections = map[int]Pair{
	1: {Left: "seed-to-soil map:\n", Right: &seedToSoil},
	2: {Left: "soil-to-fertilizer map:\n", Right: &soilToFertilizer},
	3: {Left: "fertilizer-to-water map:\n", Right: &fertilizerToWater},
	4: {Left: "water-to-light map:\n", Right: &waterToLight},
	5: {Left: "light-to-temperature map:\n", Right: &lightToTemperature},
	6: {Left: "temperature-to-humidity map:\n", Right: &temperatureToHumidity},
	7: {Left: "humidity-to-location map:\n", Right: &humidityToLocation},
}

func Part1(puzzleInput string) (solution int) {
	inputParts := strings.Split(puzzleInput, "\n\n")
	populate(inputParts)

	solution = math.MaxInt
	for _, seed := range seeds {
		track := seed

		// following the track in each section
		for i := 1; i <= len(sections); i++ {
			section := sections[i]

			found := 0
			for _, mapping := range *section.Right {
				// if the searched number is between source and source+range we found a mapping
				if track >= mapping[1] && track <= mapping[1]+mapping[2] {
					// getting the mapping destination with the diff
					found = mapping[0] + (track - mapping[1])
				}
			}

			// if its not found than its unmapped - we can keep the current value
			if found != 0 {
				track = found
			}
		}

		if solution > track {
			solution = track
		}
	}

	return solution
}

func Part2(puzzleInput string) (solution int) {
	inputParts := strings.Split(puzzleInput, "\n\n")
	populate(inputParts)

	solution = math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		//for _, seed := range seeds {
		for seed := seeds[i]; seed <= seeds[i]+seeds[i+1]; seed++ {
			fmt.Println(seed)
			track := seed

			for i := 1; i <= len(sections); i++ {
				section := sections[i]

				found := 0
				for _, mapping := range *section.Right {
					if track >= mapping[1] && track <= mapping[1]+mapping[2] {
						found = mapping[0] + (track - mapping[1])
					}
				}

				if found != 0 {
					track = found
				}
			}

			if solution > track {
				solution = track
			}
		}
	}

	return solution
}

func populate(inputParts []string) {
	if populationDone {
		return
	}

	for i, part := range inputParts {
		if i == 0 {
			for _, num := range strings.Split(strings.Replace(part, "seeds: ", "", 1), " ") {
				n, _ := strconv.Atoi(num)
				seeds = append(seeds, n)
			}
		} else {
			sectionMap := sections[i]
			header := sectionMap.Left

			var nums [][]int
			for _, line := range strings.Split(strings.Replace(part, header, "", 1), "\n") {
				parts := strings.Split(line, " ")
				source, _ := strconv.Atoi(parts[1])
				destination, _ := strconv.Atoi(parts[0])
				rnge, _ := strconv.Atoi(parts[2])

				nums = append(nums, []int{destination, source, rnge})
			}
			sections[i] = Pair{Left: header, Right: &nums}
		}
	}

	populationDone = true
}
