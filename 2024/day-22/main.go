package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-22/input.txt")
	utils.Output(day22(i, p))
}

const PRUNE_NUM = 16777216

func day22(input string, part int) int {
	nums := parseInput(input)

	if part == 1 {
		sum := 0
		for _, num := range nums {
			for range 2000 {
				num = getSecretNum(num)
			}
			sum += num
		}
		return sum
	}

	sequences := make(map[[4]int]int)
	for _, num := range nums {
		seen := make(map[[4]int]bool)
		last4 := [4]int{10, 10, 10, 10} // start with impossible changes

		for i := 0; i < 2000; i++ {
			prev := num % 10
			num = getSecretNum(num)
			curr := num % 10

			// shift the window and add new change
			last4[0] = last4[1]
			last4[1] = last4[2]
			last4[2] = last4[3]
			last4[3] = curr - prev

			// if we haven't seen this sequence before for this number
			if !seen[last4] {
				seen[last4] = true
				sequences[last4] += curr
			}
		}
	}

	// find max value in sequences map
	maxBananas := 0
	for _, sum := range sequences {
		if sum > maxBananas {
			maxBananas = sum
		}
	}
	return maxBananas
}

func parseInput(input string) []int {
	l := strings.Split(input, "\n")
	ints := make([]int, len(l))
	for i, line := range l {
		ints[i] = utils.AtoiNoErr(line)
	}
	return ints
}

func getSecretNum(num int) int {
	num = prune(mix(num*64, num))   // Step 1
	num = prune(mix(num/32, num))   // Step 2
	num = prune(mix(num*2048, num)) // Step 3
	return num
}

// -- Helpers --
func mix(a, b int) int { return a ^ b }
func prune(a int) int  { return a % PRUNE_NUM }
