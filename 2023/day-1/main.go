package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2023/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2023/day-1/input.txt")
	utils.Output(day1(i, p))
}

func day1(input string, part int) int {
	out := 0
	for _, line := range strings.Split(input, "\n") {
		nums := utils.GetIntsInString(line)
		if len(nums) == 1 {
			fmt.Printf("line: %s, nums: %v\n", line, nums)
			out += nums[0]
		} else {
			// get the first digit from the first number and the last digit from the last number
			n1 := nums[0]
			n2 := nums[len(nums)-1]
			d1 := string(strconv.Itoa(n1)[0])
			d2 := string(strconv.Itoa(n2)[len(strconv.Itoa(n2))-1])
			fmt.Printf("line: %s, d1: %s, d2: %s, combined: %s\n", line, d1, d2, d1+d2)
			out += utils.AtoiNoErr(d1 + d2)
		}
	}
	return out
}
