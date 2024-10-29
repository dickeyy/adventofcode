package main

import "github.com/dickeyy/adventofcode/2015/utils"

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./input.txt")
	utils.Output(day11(data, part))
}

func day11(data string, part int) string {
	return ""
}
