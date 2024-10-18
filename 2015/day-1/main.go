package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./data.txt")
	utils.Output(notQuiteLisp(data, part))
}

func notQuiteLisp(data string, part int) int {
	split := strings.Split(data, "")

	out := 0
	floor := 0

	for i := 0; i < len(split); i++ {
		// part 1
		if split[i] == "(" {
			floor++
		} else if split[i] == ")" {
			floor--
		}

		// part 2
		if part == 2 && floor == -1 {
			out = i + 1
			break
		}

	}

	if part == 1 {
		return floor
	}

	return out
}
