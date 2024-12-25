package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-25/input.txt")
	utils.Output(day25(i, p))
}

func day25(input string, part int) int {
	if part != 1 {
		return -1
	}

	locks, keys := parseInput(input)
	matches := 0

	for _, lock := range locks {
		for _, key := range keys {
			if checkMatch(lock, key) {
				matches++
			}
		}
	}

	return matches

	// there wasnt a part 2 for today since it was the last day !!
}

func parseInput(input string) ([][]int, [][]int) {
	var locks, keys [][]int
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i += 8 {
		if i+7 > len(lines) {
			break
		}

		heights := make([]int, 5)
		isLock := false

		// Count the '#' in each column
		for row := 0; row < 7; row++ {
			for col, char := range lines[i+row] {
				if char == '#' {
					heights[col]++
				}
			}
			// Check first row to determine if it's a lock
			if row == 0 && lines[i][0] == '#' {
				isLock = true
			}
		}

		// Adjust heights (remove the top/bottom row)
		for i := range heights {
			heights[i]--
		}

		if isLock {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	return locks, keys
}

func checkMatch(lock, key []int) bool {
	for i := 0; i < 5; i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}
