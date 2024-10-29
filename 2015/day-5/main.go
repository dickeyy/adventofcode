package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./input.txt")
	utils.Output(day5(data, part))
}

func day5(data string, part int) int {
	count := 0

	for _, line := range strings.Split(data, "\n") {
		if part == 1 {
			if isNicePt1(line) {
				count++
			}
		} else if part == 2 {
			if isNicePt2(line) {
				count++
			}
		}
	}

	return count
}

func isNicePt1(s string) bool {
	vowels := "aeiou"
	prohibited := [4]string{"ab", "cd", "pq", "xy"}

	vowelCount := 0
	hasDouble := false

	for i := 0; i < len(s); i++ {
		if strings.ContainsRune(vowels, rune(s[i])) {
			vowelCount++
		}

		if i > 0 && s[i] == s[i-1] {
			hasDouble = true
		}

		if i > 0 {
			for _, p := range prohibited {
				if s[i-1:i+1] == p {
					return false
				}
			}
		}
	}

	return vowelCount >= 3 && hasDouble
}

func isNicePt2(s string) bool {
	hasDouble := false
	hasGap := false

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Count(s, pair) >= 2 {
			hasDouble = true
			break
		}
	}

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			hasGap = true
			break
		}
	}

	return hasDouble && hasGap
}
