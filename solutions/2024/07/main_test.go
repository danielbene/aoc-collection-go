package main

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------

var solutionP1 = 3749
var exampleP1 = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_part1(t *testing.T) {
	if solutionP1 == 0 {
		return
	}

	fmt.Println("---------PART1---------")
	testSolution := Part1(exampleP1)
	if testSolution != solutionP1 {
		t.Fatalf("%d != %d", testSolution, solutionP1)
	}
}

// ---------------------------------------------------------

var solutionP2 = 11387
var exampleP2 = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_part2(t *testing.T) {
	if solutionP2 == 0 {
		return
	}

	fmt.Println("---------PART2---------")
	testSolution := Part2(exampleP2)
	if testSolution != solutionP2 {
		t.Fatalf("%d != %d", testSolution, solutionP2)
	}
}
