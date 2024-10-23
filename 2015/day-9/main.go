package main

import "github.com/dickeyy/adventofcode/2015/utils"

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./data.txt")
	utils.Output(day9(data, part))
}

func day9(data string, part int) int {
	return 0
}
