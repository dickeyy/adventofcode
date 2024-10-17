package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	data := utils.ReadFile("./data.txt")
	res1, res2 := notQuiteLisp(data)
	println(res1, res2)
}

func notQuiteLisp(data string) (int, int) {
	split := strings.Split(data, "")

	floor := 0
	posOfBasement := 0

	for i := 0; i < len(split); i++ {
		// part 1
		if split[i] == "(" {
			floor++
		} else if split[i] == ")" {
			floor--
		}

		// part 2
		if floor == -1 {
			posOfBasement = i + 1
			break
		}

	}

	return floor, posOfBasement
}
