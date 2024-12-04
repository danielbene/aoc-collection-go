package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type ProcessedInput struct {
	Lines     []string
	LineCount int
	CharCount int // length of the first line - use with care
}

func ProcessInput(puzzleInput string, pi *ProcessedInput) {
	if len(pi.Lines) != 0 {
		fmt.Println("Input already processed.")
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		line := scanner.Text()
		pi.Lines = append(pi.Lines, line)
	}

	pi.LineCount = len(pi.Lines)
	pi.CharCount = len(pi.Lines[0])
}

func Solve(solution any, duration time.Duration, name string, dir string) {
	path := dir + "solution.txt"
	f, oErr := os.OpenFile(path, os.O_RDWR, 0644)
	if oErr != nil {
		log.Fatalf("Failed to create solution file! Err: %s", oErr)
	}

	content, rErr := os.ReadFile(path)
	if rErr != nil {
		log.Fatalf("Failed to read solution file! Err: %s", oErr)
	}

	strContent := string(content)

	if name == "Part1" {
		m := regexp.MustCompile("Part1 solution:.*\n")
		strContent = m.ReplaceAllString(strContent, fmt.Sprintf("Part1 solution: %v\n", solution))

		m = regexp.MustCompile("Part1 time:.*\n")
		strContent = m.ReplaceAllString(strContent, fmt.Sprintf("Part1 time: %s\n", duration.String()))
	} else {
		m := regexp.MustCompile("Part2 solution:.*\n")
		strContent = m.ReplaceAllString(strContent, fmt.Sprintf("Part2 solution: %v\n", solution))

		m = regexp.MustCompile("Part2 time:.*\n*")
		strContent = m.ReplaceAllString(strContent, fmt.Sprintf("Part2 time: %s\n", duration.String()))
	}

	f.WriteString(strContent)

	fmt.Println("--------SOLUTIONS--------")
	fmt.Println(strContent)

	f.Close()
}
