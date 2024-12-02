package main

import (
	"aoc/util"
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
		util.Solve(p1, time.Since(start), "Part1", os.Args[2])
	} else {
		fmt.Println("running part2")
		p2 := Part2(input)
		util.Solve(p2, time.Since(start), "Part2", os.Args[2])
	}
}

// -----------------------------------------------------------

func Part1(puzzleInput string) string {
	// 12 red, 13 green, 14 blue
	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		rounds := strings.Split(parts[1], ";")

		isValid := true
		for _, round := range rounds {
			if !isValid {
				break
			}

			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				if !isValid {
					break
				}

				c := strings.Split(strings.TrimSpace(cube), " ")
				cnt, _ := strconv.Atoi(c[0])
				limit := 0

				switch c[1] {
				case "red":
					limit = 12
				case "green":
					limit = 13
				case "blue":
					limit = 14
				default:
					panic("No go.")
				}

				if cnt > limit {
					isValid = false
				}
			}
		}

		if isValid {
			id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
			sum += id
		}
	}

	return strconv.Itoa(sum)
}

func Part2(puzzleInput string) string {
	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		rounds := strings.Split(parts[1], ";")

		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				c := strings.Split(strings.TrimSpace(cube), " ")
				cnt, _ := strconv.Atoi(c[0])

				switch c[1] {
				case "red":
					if maxRed < cnt {
						maxRed = cnt
					}
				case "green":
					if maxGreen < cnt {
						maxGreen = cnt
					}
				case "blue":
					if maxBlue < cnt {
						maxBlue = cnt
					}
				default:
					panic("No go.")
				}
			}
		}

		gamePower := maxRed * maxGreen * maxBlue
		sum += gamePower
	}

	return strconv.Itoa(sum)
}
