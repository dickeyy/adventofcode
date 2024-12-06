package main

import (
    "github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
    utils.ParseFlags()
    p := utils.GetPart()

    i := utils.ReadFile("../../inputs/2015/day-19/input.txt")
    utils.Output(day19(i, p))
}

func day19(input string, part int) int {
    return 0
}

