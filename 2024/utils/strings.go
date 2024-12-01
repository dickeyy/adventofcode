package utils

import (
	"regexp"
	"strconv"
)

func GetIntsInString(s string) []int {
	re := regexp.MustCompile(`-*\d+`)
	matches := re.FindAllString(s, -1)
	ints := make([]int, len(matches))
	for i, match := range matches {
		ints[i], _ = strconv.Atoi(match)
	}
	return ints
}
