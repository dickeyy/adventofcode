package main

import (
	"regexp"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-11/input.txt")
	utils.Output(day11(i, p))
}

func day11(input string, part int) string {
	pw := input
	for !isValid(pw) {
		pw = increment(pw)
	}

	if part == 1 {
		return pw
	}

	pw = increment(pw)
	for !isValid(pw) {
		pw = increment(pw)
	}

	return pw
}

func isValid(s string) bool {
	return hasThreeLetterStraight(s) && hasNoForbiddenLetters(s) && hasTwoDifferentPairs(s)
}

func increment(s string) string {
	b := []byte(s)
	last := len(b) - 1

	i := last
	for i >= 0 {
		if b[i] == 'z' {
			b[i] = 'a'
			i--
		} else {
			b[i]++
			break
		}
	}

	return string(b)
}

func hasThreeLetterStraight(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2]+1 == s[i-1] && s[i-1]+1 == s[i] {
			return true
		}
	}
	return false
}

func hasNoForbiddenLetters(s string) bool {
	return !regexp.MustCompile("[iol]").MatchString(s)
}

func hasTwoDifferentPairs(s string) bool {
	pairs := map[string]bool{}
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			pairs[s[i-1:i+1]] = true
		}
	}
	return len(pairs) >= 2
}
