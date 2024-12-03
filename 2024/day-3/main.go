package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.txt")
	utils.Output(day3(i, p))
}

func day3(input string, part int) int {
	muls := getMulNums(input, part == 2)
	products := getProducts(muls)
	return sumProducts(products)
}

func getMulNums(input string, conditional bool) [][]int {
	// `mul\((\d+),(\d+)\)` <- this is the regex for the mul pattern matching
	// `(mul\((\d+),(\d+)\)|do\(\)|dont\(\))` <- this is the conditional regex pattern
	muls := make([][]int, 0)

	var re *regexp.Regexp
	if conditional {
		re = regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	} else {
		re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	}

	matches := re.FindAllStringSubmatchIndex(input, -1)

	allowMul := true

	for _, match := range matches {
		op := input[match[0]:match[1]]

		if conditional {
			if op == "do()" {
				allowMul = true
			} else if op == "don't()" {
				allowMul = false
			}

			if allowMul && strings.HasPrefix(op, "mul") {
				muls = append(muls, extractNumsFromMul(op))
			}
		} else {
			muls = append(muls, extractNumsFromMul(op))
		}
	}

	return muls
}

func extractNumsFromMul(op string) []int {
	numRe := regexp.MustCompile(`(\d+)`)
	nums := numRe.FindAllStringSubmatch(op, -1)
	x, _ := strconv.Atoi(nums[0][1])
	y, _ := strconv.Atoi(nums[1][1])
	return []int{x, y}
}

func getProducts(muls [][]int) []int {
	products := make([]int, 0)
	for _, mul := range muls {
		x := mul[0]
		y := mul[1]
		products = append(products, x*y)
	}
	return products
}

func sumProducts(products []int) int {
	sum := 0
	for _, product := range products {
		sum += product
	}
	return sum
}
