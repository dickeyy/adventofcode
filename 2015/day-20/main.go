package main

import (
	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-20/input.txt")
	utils.Output(day20(i, p))
}

func day20(input string, part int) int {
	target := utils.AtoiNoErr(input)

	// estimate upper bound
	lim := target / 10
	if part == 2 {
		lim = target / 10 * 2
	}
	houses := make([]int, lim+1)

	// for each elf
	for elf := 1; elf <= lim; elf++ {
		if part == 1 {
			// add this elf's contribution to all houses they visit
			for house := elf; house <= lim; house += elf {
				houses[house] += elf * 10
			}
		} else {
			// each elf visits at most 50 houses
			for mul := 1; mul <= 50 && elf*mul <= lim; mul++ {
				house := elf * mul
				houses[house] += elf * 11
			}
		}
	}

	// find the lowest house >= target
	for house := 1; house <= lim; house++ {
		if houses[house] >= target {
			return house
		}
	}

	return -1
}
