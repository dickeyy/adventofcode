package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-19/input.txt")
	utils.Output(day19(i, p))
}

func day19(input string, part int) int {
	parts := strings.Split(input, "\n\n")
	replacements := strings.Split(strings.TrimSpace(parts[0]), "\n")
	molecule := strings.TrimSpace(parts[1])

	if part == 1 {
		unique := make(map[string]bool)

		for _, r := range replacements {
			p := strings.Split(r, " => ")
			from := p[0]
			to := p[1]
			for i := 0; i <= len(molecule)-len(from); i++ {
				if molecule[i:i+len(from)] == from {
					new := molecule[:i] + to + molecule[i+len(from):]
					unique[new] = true
				}
			}
		}

		return len(unique)
	}

	// part 2
	// greedy approach - always apply longest possible reverse replacement
	steps := 0
	curr := molecule

	for curr != "e" {
		// find longest to string that appears in the curr
		bestIdx := -1
		bestLen := 0
		bestFrom := ""

		for _, r := range replacements {
			p := strings.Split(r, " => ")
			from := p[0]
			to := p[1]

			// try to find to in curr
			for i := 0; i <= len(curr)-len(to); i++ {
				if curr[i:i+len(to)] == to {
					// prefer longer replacements, and if same, prefer earlier
					if len(to) > bestLen || (len(to) == bestLen && i < bestIdx) {
						bestIdx = i
						bestLen = len(to)
						bestFrom = from
					}
				}
			}
		}

		if bestIdx == -1 {
			return -1
		}

		// apply the best replacement
		curr = curr[:bestIdx] + bestFrom + curr[bestIdx+bestLen:]
		steps++
	}

	return steps
}
