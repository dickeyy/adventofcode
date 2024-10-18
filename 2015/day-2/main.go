package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./data.txt")
	utils.Output(iWTTWBNM(data, part))
}

func iWTTWBNM(data string, part int) int {
	split := strings.Split(data, "\n")
	out := 0

	for i := 0; i < len(split); i++ {
		// we now have a string like "3x11x24". We want to split this into an array [3,11,24]
		split2 := strings.Split(split[i], "x")
		// we now have an array ["3","11","24"]. We want to convert each element to an int
		ints := make([]int, len(split2))
		for j := 0; j < len(split2); j++ {
			ints[j], _ = strconv.Atoi(split2[j])
		}

		// part 1
		if part == 1 {
			sides := []int{(ints[0] * ints[1]), (ints[1] * ints[2]), (ints[2] * ints[0])}

			min := sides[0]
			for i := 1; i < len(sides); i++ {
				if sides[i] < min {
					min = sides[i]
				}
			}

			out += (2 * sides[0]) + (2 * sides[1]) + (2 * sides[2]) + min
		}

		// part 2
		if part == 2 {
			// sort the array smallest to largest
			sort.Ints(ints)
			out += ints[0] + ints[0] + ints[1] + ints[1] + (ints[0] * ints[1] * ints[2])
		}

	}

	return out
}
