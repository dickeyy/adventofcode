package main

import (
	"slices"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./data.txt")
	res := pshiav(data, part)
	println(res)
}

// ^ = north, > = east, v = south, < = west
func pshiav(data string, part int) int {
	vc := make([][]int, 0) // houses we have already visited

	cc1 := make([]int, 2) // current coordinate santa 1
	cc2 := make([]int, 2) // current coordinate santa 2

	// add the starting location
	vc = append(vc, []int{0, 0})

	cs := 0 // current santa -> 0 = santa 1, 1 = santa 2

	for _, line := range strings.Split(data, "") {
		var currentCC []int
		if cs == 0 {
			cc1 = moveSanta(cc1, line)
			currentCC = cc1
		} else if cs == 1 && part == 2 {
			cc2 = moveSanta(cc2, line)
			currentCC = cc2
		}

		if !alreadyVisted(vc, currentCC) {
			vc = append(vc, append([]int(nil), currentCC...))
		}

		if part == 2 {
			cs = 1 - cs
		}
	}

	return len(vc)
}

// helper func to check if a coordinate has been visited before
func alreadyVisted(vc [][]int, cc []int) bool {
	return slices.ContainsFunc(vc, func(v []int) bool {
		return v[0] == cc[0] && v[1] == cc[1]
	})
}

// helper func to handle movement for a given santa
func moveSanta(cc []int, dir string) []int {
	switch dir {
	case "^":
		cc[1]++
	case ">":
		cc[0]++
	case "v":
		cc[1]--
	case "<":
		cc[0]--
	}
	return cc
}
