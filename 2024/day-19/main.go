package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-19/input.txt")
	utils.Output(day19(i, p))
}

func day19(input string, part int) int {
	patterns, designs := parseInput(input)
	out := 0

	// keep track of all strings that have been seen
	cache := make(map[string]int)

	// recursive function to solve the problem (defined in scope for access to other variables)
	var solve func(string) int
	solve = func(s string) int {
		// if s is not in cache
		if _, ok := cache[s]; !ok {
			if len(s) == 0 {
				return 1
			}
			res := 0
			// check each pattern to see if it matches (recursively)
			for _, pattern := range patterns {
				if strings.HasPrefix(s, string(pattern)) {
					res += solve(s[len(pattern):])
				}
			}
			// cache the result
			cache[s] = res
		}
		// return the cached result (guaranteed to exist)
		return cache[s]
	}

	for _, design := range designs {
		if part == 1 {
			if solve(design) > 0 {
				out++
			}
		} else {
			out += solve(design)
		}
	}

	return out
}

func parseInput(input string) ([]string, []string) {
	s := strings.Split(input, "\n\n")
	return strings.Split(strings.TrimSpace(s[0]), ", "), strings.Split(strings.TrimSpace(s[1]), "\n")
}
