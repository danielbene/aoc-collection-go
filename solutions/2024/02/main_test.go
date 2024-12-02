package main

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------

var solutionP1 = `2`
var exampleP1 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_part1(t *testing.T) {
	fmt.Println("---------PART1---------")
	testSolution := Part1(exampleP1)
	if testSolution != solutionP1 {
		t.Fatalf("%s != %s", testSolution, solutionP1)
	}
}

// ---------------------------------------------------------

var solutionP2 = `4`
var exampleP2 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_part2(t *testing.T) {
	fmt.Println("---------PART2---------")
	testSolution := Part2(exampleP2)
	if testSolution != solutionP2 {
		t.Fatalf("%s != %s", testSolution, solutionP2)
	}
}
