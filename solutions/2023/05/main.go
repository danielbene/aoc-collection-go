package main

import (
	"aoc/util/aocutil"
	"aoc/util/pairutil"
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

func Part1(puzzleInput string) (solution int) {
	inputParts := strings.Split(puzzleInput, "\n\n")

	var (
		seeds                 []int
		seedToSoil            pairutil.Pairs
		soilToFertilizer      pairutil.Pairs
		fertilizerToWater     pairutil.Pairs
		waterToLight          pairutil.Pairs
		lightToTemperature    pairutil.Pairs
		temperatureToHumidity pairutil.Pairs
		humidityToLocation    pairutil.Pairs
	)

	var sections = map[int]pairutil.Pair{
		1: {Left: "seed-to-soil map:\n", Right: &seedToSoil},
		2: {Left: "soil-to-fertilizer map:\n", Right: &soilToFertilizer},
		3: {Left: "fertilizer-to-water map:\n", Right: &fertilizerToWater},
		4: {Left: "water-to-light map:\n", Right: &waterToLight},
		5: {Left: "light-to-temperature map:\n", Right: &lightToTemperature},
		6: {Left: "temperature-to-humidity map:\n", Right: &temperatureToHumidity},
		7: {Left: "humidity-to-location map:\n", Right: &humidityToLocation},
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
			pairs := sectionMap.Right

			for _, line := range strings.Split(strings.Replace(part, header.(string), "", 1), "\n") {
				parts := strings.Split(line, " ")
				source, _ := strconv.Atoi(parts[1])
				destination, _ := strconv.Atoi(parts[0])
				rnge, _ := strconv.Atoi(parts[2])

				for i := 0; i < rnge; i++ {
					// OMFG
					(pairs.(*pairutil.Pairs)).Pairs = append((pairs.(*pairutil.Pairs)).Pairs, pairutil.Pair{Left: source + i, Right: destination + i})
				}
			}
		}
	}

	solution = math.MaxInt
	for _, seed := range seeds {
		track := seed
		for i := 1; i <= len(sections); i++ {
			section := sections[i]
			tmp, found := (section.Right.(*pairutil.Pairs)).GetFirstByLeft(track)

			// if its not found than its unmapped - we can keep the current value
			if found {
				track = tmp.(int)
			}
		}

		if solution > track {
			solution = track
		}
	}

	return solution
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	fmt.Printf("%d - %d\n", input.LineCount, input.CharCount)

	return solution
}
