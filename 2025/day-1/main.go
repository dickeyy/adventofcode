package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2025/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2025/day-1/input.txt")
	utils.Output(day1(i, p))
}

func countZeros(lo, hi int) int {
	if lo > hi {
		return 0
	}
	return utils.FloorDiv(hi, 100) - utils.FloorDiv(lo-1, 100)
}

func day1(input string, part int) int {
	currPos := 50
	numAtZero := 0

	for line := range strings.SplitSeq(input, "\n") {
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])

		switch dir {
		case 'L':
			if part == 2 {
				numAtZero += countZeros(currPos-steps, currPos-1)
			}
			currPos = (currPos - steps) % 100
		case 'R':
			if part == 2 {
				numAtZero += countZeros(currPos+1, currPos+steps)
			}
			currPos = (currPos + steps) % 100
		}

		if part == 1 {
			if currPos == 0 {
				numAtZero++
			}
		}
	}

	return numAtZero
}
