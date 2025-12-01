package utils

import (
	"regexp"
	"strconv"
	"strings"
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

func AtoiNoErr(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func IntArrayToString(arr []int, sep string) string {
	strNums := make([]string, len(arr))
	for i, num := range arr {
		strNums[i] = strconv.Itoa(num)
	}
	return strings.Join(strNums, sep)
}
