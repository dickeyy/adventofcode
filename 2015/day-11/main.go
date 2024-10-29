package main

import "github.com/dickeyy/adventofcode/2015/utils"

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	input := utils.ReadFile("./input.txt")
	utils.Output(day11(input, part))
}

func day11(input string, part int) string {
	return ""
}
