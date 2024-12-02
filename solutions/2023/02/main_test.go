package main

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------

var solutionP1 = `8`
var exampleP1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func Test_part1(t *testing.T) {
	fmt.Println("---------PART1---------")
	testSolution := Part1(exampleP1)
	if testSolution != solutionP1 {
		t.Fatalf("%s != %s", testSolution, solutionP1)
	}
}

// ---------------------------------------------------------

var solutionP2 = `2286`
var exampleP2 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func Test_part2(t *testing.T) {
	fmt.Println("---------PART2---------")
	testSolution := Part2(exampleP2)
	if testSolution != solutionP2 {
		t.Fatalf("%s != %s", testSolution, solutionP2)
	}
}
