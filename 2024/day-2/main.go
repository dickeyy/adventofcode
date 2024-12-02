package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.txt")
	utils.Output(day2(i, p))
}

func day2(input string, part int) int {
	out := 0

	matrix := parseMatrix(input)

	for _, row := range matrix {
		if part == 1 {
			if isValidRow(row) {
				out++
			}
			continue
		}

		if isValidRow(row) || canBecomeValid(row) {
			out++
		}
	}

	return out
}

func parseMatrix(input string) [][]int {
	lines := strings.Split(input, "\n")

	matrix := make([][]int, len(lines))

	for i, line := range lines {
		nums := strings.Fields(line)

		row := make([]int, len(nums))
		for j, num := range nums {
			n, _ := strconv.Atoi(num)
			row[j] = n
		}

		matrix[i] = row
	}

	return matrix
}

// checks the following conditions:
// 1. Must consistently increase or decrease (no switching)
// 2. Differences between adjacent numbers must be between 1 and 3 (inclusive)
func isValidRow(row []int) bool {
	if len(row) < 2 {
		return true
	}

	// determine the initial direction from the first 2 nums
	isInc := row[1] > row[0]

	// check first pair meets difference criteria
	initialDiff := utils.Abs(row[1] - row[0])
	if initialDiff < 1 || initialDiff > 3 {
		return false
	}

	// check remaining pairs
	for i := 1; i < len(row)-1; i++ {
		curr, next := row[i], row[i+1]
		diff := next - curr

		// if direction changes, sequence is invalid
		if (diff > 0) != isInc {
			return false
		}

		// check the diff is within bounds
		if utils.Abs(diff) < 1 || utils.Abs(diff) > 3 {
			return false
		}
	}

	return true
}

// this func checks if by removing a single number from a row, the once bad row becomes valid
func canBecomeValid(row []int) bool {
	// try removing each num and check if the resulting row is valid
	for i := range row {
		// create a new slice without the current num
		newRow := make([]int, 0, len(row)-1)
		newRow = append(newRow, row[:i]...)
		newRow = append(newRow, row[i+1:]...)

		if isValidRow(newRow) {
			return true
		}
	}
	return false
}
