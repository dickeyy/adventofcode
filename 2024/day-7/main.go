package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-7/input.txt")
	utils.Output(day7(i, p))
}

type Operator int

const (
	Add Operator = iota
	Multiply
)

func day7(input string, part int) int {
	calibrations := parseInput(input)
	validKeys := make([]int, 0)

	for key, vals := range calibrations {
		if isCalibrationValid(key, vals, part == 2) {
			validKeys = append(validKeys, key)
		}
	}

	// sum the valid keys
	return utils.SumNums(validKeys)
}

func parseInput(input string) map[int][]int {
	calibrations := make(map[int][]int)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		key, _ := strconv.Atoi(parts[0])

		for _, val := range strings.Split(parts[1], " ") {
			val, _ := strconv.Atoi(val)
			calibrations[key] = append(calibrations[key], val)
		}
	}

	return calibrations
}

func isCalibrationValid(target int, vals []int, tryConcat bool) bool {
	if len(vals) < 2 {
		return vals[0] == target
	}

	var backtrack func(pos, curr int) bool
	backtrack = func(pos, curr int) bool {
		// if we've used all the numbers, check if we hit the target
		if pos == len(vals) {
			return curr == target
		}

		// try addition
		if backtrack(pos+1, curr+vals[pos]) {
			return true
		}

		// try multiplication
		if backtrack(pos+1, curr*vals[pos]) {
			return true
		}

		if tryConcat {
			// convert the current to a string, concat the next, convert back
			currStr := strconv.Itoa(curr)
			nextStr := strconv.Itoa(vals[pos])
			concatNum, _ := strconv.Atoi(currStr + nextStr)
			if backtrack(pos+1, concatNum) {
				return true
			}
		}

		return false
	}

	return backtrack(1, vals[0])
}
