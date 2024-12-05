package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-17/input.txt")
	utils.Output(day17(i, p))
}

type Container struct {
	capacity int
}

const MAX_CAPACITY = 150

func day17(input string, part int) int {
	containers := make([]Container, 0)

	for _, line := range strings.Split(input, "\n") {
		c := Container{}
		c.capacity, _ = strconv.Atoi(line)
		containers = append(containers, c)
	}

	if part == 1 {
		return findCombinations(containers, MAX_CAPACITY)
	}

	return findMinCombinations(containers, MAX_CAPACITY)
}

func findCombinations(containers []Container, target int) int {
	count := 0
	var backtrack func(inex, currentSum int)

	backtrack = func(index, currentSum int) {
		// found a valid combination
		if currentSum == target {
			count++
			return
		}

		// gone over target or reached end of containers
		if currentSum > target || index >= len(containers) {
			return
		}

		// skip this container
		backtrack(index+1, currentSum)

		// use this container
		backtrack(index+1, currentSum+containers[index].capacity)
	}

	backtrack(0, 0)
	return count
}

func findMinCombinations(containers []Container, target int) int {
	minContainers := len(containers) + 1
	countWithMin := 0

	var backtrack func(index, currentSum, containersUsed int)

	backtrack = func(index, currentSum, containersUsed int) {
		// found a valid combination
		if currentSum == target {
			if containersUsed < minContainers {
				// found a new min
				minContainers = containersUsed
				countWithMin = 1
			} else if containersUsed == minContainers {
				// found another combo with the same min
				countWithMin++
			}
			return
		}

		// gone over target or reached end of containers
		if currentSum > target || index >= len(containers) {
			return
		}

		// skip this container
		backtrack(index+1, currentSum, containersUsed)

		// use this container
		backtrack(index+1, currentSum+containers[index].capacity, containersUsed+1)
	}

	backtrack(0, 0, 0)
	return countWithMin
}
