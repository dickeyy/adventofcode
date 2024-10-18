package main

import "github.com/dickeyy/adventofcode/2015/utils"

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./data.txt")
	res := tiss(data, part)
	println("Output: ", res)
}

func tiss(data string, part int) int {
	return 0
}
