package main

import (
    "github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
    utils.ParseFlags()
    p := utils.GetPart()

    i := utils.ReadFile("../../inputs/2024/day-6/input.txt")
    utils.Output(day6(i, p))
}

func day6(input string, part int) int {
    return 0
}

