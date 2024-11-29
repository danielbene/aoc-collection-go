package main

import (
	"aoc/util"
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
		p1 := part1()
		util.Solve(p1, time.Since(start), "Part1", os.Args[2])
	} else {
		p2 := part2()
		util.Solve(p2, time.Since(start), "Part2", os.Args[2])
	}
}

func part1() string {
	fmt.Println("running part1")
	return "solution"
}

func part2() string {
	fmt.Println("running part2")
	return "solution"
}
