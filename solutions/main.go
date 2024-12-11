package main

import (
	"aoc/util/matrixutil"
	"fmt"
)

func main() {
	var b [][]string
	var c [][]string
	var d [][]string
	a := [][]string{{"asd", "vbn"}, {"qwe"}}
	d = a

	b = matrixutil.DeepCopy(a)
	c = matrixutil.DeepCopy(a)

	b = append(b, []string{"fgh"})
	c = append(c, []string{"fgh", "qwe"})

	c = b
	b = append(b, []string{"ttt"})

	a = append(a, []string{"aaaa"})

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
