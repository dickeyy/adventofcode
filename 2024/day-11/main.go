package main

import (
	"strconv"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-11/input.txt")
	utils.Output(day11(i, p))
}

func day11(input string, part int) int {
	s := utils.GetIntsInString(input)

	// count freq of each stone val
	scs := make(map[int]int)
	for _, s := range s {
		scs[s]++
	}

	iters := 25
	if part == 2 {
		iters = 75
	}

	// process stones
	for i := 0; i < iters; i++ {
		scs = blinkAllStones(scs)
	}

	return sumStones(scs)
}

func blinkOnce(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	// convert to string to check digits
	s := strconv.Itoa(stone)

	if len(s)%2 == 0 {
		// split into two halves
		halfway := len(s) / 2
		l := utils.AtoiNoErr(s[:halfway])
		r := utils.AtoiNoErr(s[halfway:])
		return []int{l, r}
	}

	return []int{stone * 2024}
}

func blinkAllStones(stoneCount map[int]int) map[int]int {
	nc := make(map[int]int)

	for s, c := range stoneCount {
		// get new stones from transformation
		ns := blinkOnce(s)
		// add to counts, multiplied by how many of oritinal stones we had
		for _, n := range ns {
			nc[n] += c
		}
	}

	return nc
}

func sumStones(stoneCount map[int]int) int {
	total := 0
	for _, c := range stoneCount {
		total += c
	}
	return total
}
