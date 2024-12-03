package main

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------

var solutionP1 = `4361`
var exampleP1 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_part1(t *testing.T) {
	fmt.Println("---------PART1---------")
	testSolution := Part1(exampleP1)
	if testSolution != solutionP1 {
		t.Fatalf("%s != %s", testSolution, solutionP1)
	}
}

// ---------------------------------------------------------

var solutionP2 = `467835`
var exampleP2 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_part2(t *testing.T) {
	fmt.Println("---------PART2---------")
	testSolution := Part2(exampleP2)
	if testSolution != solutionP2 {
		t.Fatalf("%s != %s", testSolution, solutionP2)
	}
}
