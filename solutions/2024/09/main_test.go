package main

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------

var solutionP1 = 1928
var exampleP1 = `2333133121414131402`

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

var solutionP2 = 0
var exampleP2 = `asd
qwe
fgh`

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
