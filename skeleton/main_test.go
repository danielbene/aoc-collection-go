package main

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------

var solutionP1 = `x`
var exampleP1 = `asd
qwe
fgh`

func Test_part1(t *testing.T) {
	fmt.Println("---------PART1---------")
	testSolution := Part1(exampleP1)
	if testSolution != solutionP1 {
		t.Fatalf("%s != %s", testSolution, solutionP1)
	}
}

// ---------------------------------------------------------

var solutionP2 = `x`
var exampleP2 = `asd
qwe
fgh`

func Test_part2(t *testing.T) {
	return
	fmt.Println("---------PART2---------")
	testSolution := Part2(exampleP2)
	if testSolution != solutionP2 {
		t.Fatalf("%s != %s", testSolution, solutionP2)
	}
}
