package main

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------

var solutionP1 = 14
var exampleP1 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

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
