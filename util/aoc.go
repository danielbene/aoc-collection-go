package util

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func Solve(solution string, duration time.Duration, name string, dir string) {
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
		strContent = m.ReplaceAllString(strContent, fmt.Sprintf("Part1 solution: %s\n", solution))

		m = regexp.MustCompile("Part1 time:.*\n")
		strContent = m.ReplaceAllString(strContent, fmt.Sprintf("Part1 time: %s\n", duration.String()))
	} else {
		m := regexp.MustCompile("Part2 solution:.*\n")
		strContent = m.ReplaceAllString(strContent, fmt.Sprintf("Part2 solution: %s\n", solution))

		m = regexp.MustCompile("Part2 time:.*\n*")
		strContent = m.ReplaceAllString(strContent, fmt.Sprintf("Part2 time: %s\n", duration.String()))
	}

	f.WriteString(strContent)

	fmt.Println("--------SOLUTIONS--------")
	fmt.Println(strContent)

	f.Close()
}
