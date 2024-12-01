package main

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------

var solutionP1 = `11`
var exampleP1 = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test_part1(t *testing.T) {
	fmt.Println("---------PART1---------")
	testSolution := Part1(exampleP1)
	if testSolution != solutionP1 {
		t.Fatalf("%s != %s", testSolution, solutionP1)
	}
}

// ---------------------------------------------------------

var solutionP2 = `31`
var exampleP2 = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test_part2(t *testing.T) {
	fmt.Println("---------PART2---------")
	testSolution := Part2(exampleP2)
	if testSolution != solutionP2 {
		t.Fatalf("%s != %s", testSolution, solutionP2)
	}
}
