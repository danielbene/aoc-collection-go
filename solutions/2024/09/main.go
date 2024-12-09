package main

import (
	"aoc/util/aocutil"
	"aoc/util/sliceutil"
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputFile string

func init() {
	inputFile = strings.TrimRight(inputFile, "\n")
}

func main() {
	start := time.Now()
	if os.Args[1] == "1" {
		fmt.Println("running part1")
		p1 := Part1(inputFile)
		aocutil.Solve(p1, time.Since(start), "Part1", os.Args[2])
	} else {
		fmt.Println("running part2")
		p2 := Part2(inputFile)
		aocutil.Solve(p2, time.Since(start), "Part2", os.Args[2])
	}
}

// -----------------------------------------------------------

var input aocutil.ProcessedInput

func Part1(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	var diskLayout []int

	blockId := 0
	for _, line := range input.Lines {
		for idx, ch := range line {
			num, _ := strconv.Atoi(string(ch))
			if idx%2 == 0 {
				for i := 0; i < num; i++ {
					diskLayout = append(diskLayout, blockId)
				}

				blockId++
			} else {
				for i := 0; i < num; i++ {
					diskLayout = append(diskLayout, -1) // using -1 as point
				}
			}
		}
	}

	for i := len(diskLayout) - 1; i >= 0; i-- {
		if diskLayout[i] == -1 {
			continue
		}

		for j := 0; j < len(diskLayout); j++ {
			if j > i {
				// if forward index reaches backward we should stop
				break
			}

			if diskLayout[j] == -1 {
				diskLayout[j], diskLayout[i] = diskLayout[i], diskLayout[j]
			}

		}
	}

	for idx, num := range diskLayout {
		if num == -1 {
			break
		}

		solution += idx * num
	}

	return solution
}

func Part2(puzzleInput string) (solution int) {
	aocutil.ProcessInput(puzzleInput, &input)

	var diskLayout [][]int

	blockId := 0
	for _, line := range input.Lines {
		for idx, ch := range line {
			num, _ := strconv.Atoi(string(ch))
			if idx%2 == 0 {
				var file []int
				for i := 0; i < num; i++ {
					file = append(file, blockId)
				}

				diskLayout = append(diskLayout, file)

				blockId++
			} else {
				for i := 0; i < num; i++ {
					diskLayout = append(diskLayout, []int{-1}) // using -1 as point
				}
			}
		}
	}

	// using index saving and goto label to avoid issues with slice element manipulation inside loop
	saveIndex := len(diskLayout) - 1

MAINLOOP:
	// backward loop that searching for movable blocks
	for idxBackward := saveIndex; idxBackward >= 0; idxBackward-- {
		if diskLayout[idxBackward][0] == -1 {
			continue
		}

		// forward loop that searching for free spaces
		for idxForward := 0; idxForward < len(diskLayout); idxForward++ {
			if idxForward > idxBackward {
				// if forward index reaches backward we should stop
				break
			}

			if diskLayout[idxForward][0] == -1 {
				lenNum := len(diskLayout[idxBackward])
				fits := true

				// free space found, check forward if there is enough space for the block
				for k := 0; k < lenNum; k++ {
					if diskLayout[idxForward+k][0] != -1 {
						fits = false
						break
					}
				}

				if fits {
					// swap free space and file block
					diskLayout[idxForward], diskLayout[idxBackward] = diskLayout[idxBackward], diskLayout[idxForward]

					// move free space to the file block index while block elemnt count reached
					for k := 1; k < lenNum; k++ {
						// using only +1 in the indexes because of the continous element removal
						diskLayout = slices.Insert(diskLayout, idxBackward+1, []int{-1})
						diskLayout = sliceutil.RemoveIntArrSliceElement(diskLayout, idxForward+1)
					}

					saveIndex = idxBackward - 1
					goto MAINLOOP
				}
			}
		}
	}

	// TODO: add flatten to matrixutil
	var flatten []int
	for _, outArr := range diskLayout {
		flatten = append(flatten, outArr...)
	}

	for idx, num := range flatten {
		if num == -1 {
			continue
		}

		solution += idx * num
	}

	return solution
}
