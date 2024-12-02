package main

import (
	"aoc/util"
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/divan/num2words"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
}

func main() {
	start := time.Now()
	if os.Args[1] == "1" {
		fmt.Println("running part1")
		p1 := Part1(input)
		util.Solve(p1, time.Since(start), "Part1", os.Args[2])
	} else {
		fmt.Println("running part2")
		p2 := Part2(input)
		util.Solve(p2, time.Since(start), "Part2", os.Args[2])
	}
}

// -----------------------------------------------------------

func Part1(puzzleInput string) string {
	sum := counter(puzzleInput)
	return fmt.Sprintf("%d", sum)
}

func Part2(puzzleInput string) string {
	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		// storing all the decimal and word formatted numbers (map value)
		// AND the index (map key) inside the line
		nums := make(map[int]int)

		line := scanner.Text()
		for idx, ch := range line {
			if unicode.IsDigit(ch) {
				decNum, _ := strconv.Atoi(string(ch))
				nums[idx] = decNum
			}
		}

		// checking only the first 20 number word
		// index has to be checked multiple times for repetitions
		for i := 0; i < 20; i++ {
			word := num2words.Convert(i)

			startIdx := 0
			for {
				idx := strings.Index(line[startIdx:], word)
				if idx != -1 {
					nums[startIdx+idx] = i
					startIdx = startIdx + idx + 1
				} else {
					break
				}
			}
		}

		keys := make([]int, 0, len(nums))
		for k := range nums {
			keys = append(keys, k)
		}

		sort.Ints(keys)

		// splitting 10+ decimals with modulo
		dec, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[keys[0]], nums[keys[len(keys)-1]]%10))
		sum += dec
	}

	return strconv.Itoa(sum)
}

func counter(puzzleInput string) int {
	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for scanner.Scan() {
		num := ""
		line := scanner.Text()

		for _, ch := range line {
			if unicode.IsDigit(ch) {
				num = string(ch)
				break
			}
		}

		runes := []rune(line)
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				num = num + string(runes[i])
				break
			}
		}

		decNum, _ := strconv.Atoi(num)
		sum += decNum
	}

	return sum
}
