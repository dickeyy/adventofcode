package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-1/input.txt")
	utils.Output(day1(i, p))
}

func day1(input string, part int) int {
	split := strings.Split(input, "")

	out := 0
	floor := 0

	for i := 0; i < len(split); i++ {
		// part 1
		if split[i] == "(" {
			floor++
		} else if split[i] == ")" {
			floor--
		}

		// part 2
		if part == 2 && floor == -1 {
			out = i + 1
			break
		}

	}

	if part == 1 {
		return floor
	}

	return out
}
