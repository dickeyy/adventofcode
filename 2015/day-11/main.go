package main

import "github.com/dickeyy/adventofcode/2015/utils"

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.txt")
	utils.Output(day11(i, p))
}

func day11(input string, part int) string {
	return ""
}
