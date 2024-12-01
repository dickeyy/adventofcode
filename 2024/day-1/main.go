package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.txt")
	utils.Output(day1(i, p))
}

func day1(input string, part int) int {
	out := 0

	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		l, _ := strconv.Atoi(strings.Split(line, "   ")[0])
		r, _ := strconv.Atoi(strings.Split(line, "   ")[1])

		left = append(left, l)
		right = append(right, r)
	}

	if part == 1 {
		// sort left and right in ascending order
		sort.Ints(left)
		sort.Ints(right)

		// for each pair of numbers, find the difference
		for i := 0; i < len(left); i++ {
			out += int(math.Abs(float64(left[i] - right[i])))
		}
	} else {
		// make a hashmap for right frequency
		freq := make(map[int]int)
		for _, num := range right {
			freq[num]++
		}

		// for each number in left, find the number of times it appears in right
		for _, num := range left {
			if f, exists := freq[num]; exists {
				out += num * f
			}
		}
	}

	return out
}
