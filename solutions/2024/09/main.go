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

	ordered := diskLayout
	for i := len(ordered) - 1; i >= 0; i-- {
		if ordered[i] == -1 {
			continue
		}

		for j := 0; j < len(ordered); j++ {
			if j > i {
				// if forward index reaches backward we should stop
				break
			}

			if ordered[j] == -1 {
				ordered[j], ordered[i] = ordered[i], ordered[j]
			}
		}

	}

	for idx, num := range ordered {
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

MAINLOOP:
	fmt.Println(diskLayout)
	for idxBackward := len(diskLayout) - 1; idxBackward >= 0; idxBackward-- {
		if diskLayout[idxBackward][0] == -1 {
			continue
		}

		for idxForward := 0; idxForward < len(diskLayout); idxForward++ {
			if idxForward > idxBackward {
				// if forward index reaches backward we should stop
				break
			}

			if diskLayout[idxForward][0] == -1 {
				// len num
				// check forward

				// TODO: handle idx > 9
				lenNum := len(diskLayout[idxBackward])

				fits := true
				for k := 0; k < lenNum; k++ {
					if diskLayout[idxForward+k][0] != -1 {
						fits = false
						break
					}
				}

				if fits {
					fmt.Printf("%v - %v\n", diskLayout[idxForward], diskLayout[idxBackward])
					diskLayout[idxForward], diskLayout[idxBackward] = diskLayout[idxBackward], diskLayout[idxForward]

					for k := 1; k < lenNum; k++ {
						diskLayout = slices.Insert(diskLayout, idxBackward+k, []int{-1})

						//FIXME: something fishy
						fmt.Printf("delete: %v\n", diskLayout[idxForward+k])
						diskLayout = sliceutil.RemoveIntArrSliceElement(diskLayout, idxForward+k)
					}

					goto MAINLOOP
				}
			}
		}
	}

	/*ordered := diskLayout
	for i := len(ordered) - 1; i >= 0; i-- {
		if ordered[i] == -1 {
			continue
		}

		for j := 0; j < len(ordered); j++ {
			if j > i {
				// if forward index reaches backward we should stop
				break
			}

			if ordered[j] == -1 {
				// len num
				// check forward

				lenNum := len(strconv.Itoa(ordered[i]))

				fits := true
				for k := 0; k < lenNum; k++ {
					if ordered[j+k] != -1 {
						fits = false
						break
					}
				}

				if fits {
					ordered[j], ordered[i] = ordered[i], ordered[j]

					for k := 1; k < lenNum; k++ {
						sliceutil.RemoveIntSliceElement(ordered, j+k)
						ordered = append(ordered, -1)
					}

					fmt.Println(ordered)
				}
			}
		}
	}

	for idx, num := range ordered {
		if num == -1 {
			break
		}

		solution += idx * num
	}*/

	return solution
}
