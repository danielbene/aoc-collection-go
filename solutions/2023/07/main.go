package main

import (
	"aoc/util/aocutil"
	_ "embed"
	"fmt"
	"os"
	"slices"
	"sort"
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

type Hand struct {
	Cards string
	Bid   int
}

var input aocutil.ProcessedInput
var cardTypes = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

func Part1(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	var ranked []Hand
	for _, line := range input.Lines {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])

		ranked = append(ranked, Hand{Cards: parts[0], Bid: bid})
	}

	sort.Slice(ranked, func(i, j int) bool {
		return compareHands(ranked[i], ranked[j])
	})

	handCnt := len(ranked)
	for i, hand := range ranked {
		solution += hand.Bid * (handCnt - i)
	}

	return solution
}

func compareHands(left Hand, right Hand) bool {
	leftScore := scoreHand(left)
	rightScore := scoreHand(right)

	if leftScore == rightScore {
		for i := 0; i < len(left.Cards); i++ {
			lCard := slices.Index(cardTypes, string(left.Cards[i]))
			rCard := slices.Index(cardTypes, string(right.Cards[i]))

			if lCard == rCard {
				continue
			} else if lCard < rCard {
				return true
			} else {
				return false
			}
		}
	} else if leftScore > rightScore {
		return true
	}

	return false
}

func scoreHand(hand Hand) int {
	score := 0
	for _, card := range cardTypes {
		cnt := strings.Count(hand.Cards, card)

		switch {
		case cnt == 2:
			score += 1
		case cnt == 3:
			score += 3
		case cnt == 4:
			score += 5
		case cnt == 5:
			score += 6
		}
	}

	return score
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	fmt.Printf("%d - %d\n", input.LineCount, input.CharCount)

	return solution
}
