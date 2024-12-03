package main

import (
    "github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
    utils.ParseFlags()
    p := utils.GetPart()

    i := utils.ReadFile("./input.txt")
    utils.Output(day15(i, p))
}

func day15(input string, part int) int {
    return 0
}

