package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./data.txt")
	utils.Output(day10(data, part))
}

func day10(data string, part int) int {
	var iterations int

	if part == 1 {
		iterations = 40
	} else {
		iterations = 50
	}

	out := lookAndSay(data, iterations)

	return len(out)
}

func lookAndSay(s string, iterations int) string {
	/*
		We are given a string of numbers
		we want to count the number number of same digit runs and describe them
		so for example:
			1 -> 11 (i see 1 one)
			11 -> 21 (i see 2 ones in a run)
			21 -> 1211 (i see 1 two, and 1 one)
			1211 -> 111221
			111221 -> 312211
			312211 -> 13112221
			and so on...
		we want to recurisvely build up a string until iterations = 0
	*/

	// base case
	if iterations == 0 {
		return s
	}

	var out strings.Builder
	currentCount := 1

	for i := 0; i < len(s); i++ {
		// if we are not at the last character and the next digit is the same as the current
		if i < len(s)-1 && s[i+1] == s[i] {
			currentCount++
		} else {
			// end of a run, write count and digit
			out.WriteString(strconv.Itoa(currentCount))
			out.WriteByte(s[i])
			currentCount = 1
		}
	}

	// recurse
	return lookAndSay(out.String(), iterations-1)
}
