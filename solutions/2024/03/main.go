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
	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		for _, val := range r.FindAllStringSubmatch(line, -1) {
			a, _ := strconv.Atoi(val[1])
			b, _ := strconv.Atoi(val[2])
			sum += a * b
		}
	}

	return strconv.Itoa(sum)
}

func Part2(puzzleInput string) string {
	sum := 0
	skip := false

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
		for _, val := range r.FindAllStringSubmatch(line, -1) {
			if val[0] == "do()" {
				skip = false
			} else if val[0] == "don't()" {
				skip = true
			}

			if !skip {
				a, _ := strconv.Atoi(val[1])
				b, _ := strconv.Atoi(val[2])
				sum += a * b
			}
		}
	}

	return strconv.Itoa(sum)
}
