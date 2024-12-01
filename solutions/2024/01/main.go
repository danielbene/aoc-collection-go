package main

import (
	"aoc/util"
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"slices"
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

var (
	arrLeft  []uint32
	arrRight []uint32
)

func Part1(puzzleInput string) string {
	processInput(puzzleInput)

	var sum uint32
	for i := 0; i < len(arrLeft); i++ {
		l := arrLeft[i]
		r := arrRight[i]

		if l <= r {
			sum += r - l
		} else {
			sum += l - r
		}
	}

	return strconv.Itoa(int(sum))
}

func Part2(puzzleInput string) string {
	processInput(puzzleInput)

	var (
		sum     uint32
		arrDone = make(map[uint32]uint32)
	)

	for i := 0; i < len(arrLeft); i++ {
		l := arrLeft[i]
		itemCnt, exists := arrDone[l]

		if exists {
			sum += l * itemCnt
			continue
		}

		for _, item := range arrRight {
			if l == item {
				itemCnt++
			}
		}

		arrDone[l] = uint32(itemCnt)
		sum += uint32(l * itemCnt)
	}

	return strconv.Itoa(int(sum))
}

func processInput(puzzleInput string) {
	arrLeft = []uint32{}
	arrRight = []uint32{}

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		arrLeft = append(arrLeft, uint32(left))
		arrRight = append(arrRight, uint32(right))
	}

	slices.Sort(arrLeft)
	slices.Sort(arrRight)
}
