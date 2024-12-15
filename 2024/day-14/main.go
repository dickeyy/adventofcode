package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-14/input.txt")
	utils.Output(day14(i, p))
}

type Coordinate struct {
	x, y int
}

type Robot struct {
	pos Coordinate // position
	vel Coordinate // velocity
}

func day14(input string, part int) int {
	robots := parseInput(input)
	if part == 1 {
		return simulateAndCalculate(robots)
	}
	return findEasterEgg(robots)
}

func parseInput(input string) []Robot {
	lines := strings.Split(input, "\n")
	robots := make([]Robot, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		// please dont judge me for the following three lines... it works...
		robots[i] = Robot{
			pos: Coordinate{utils.AtoiNoErr(strings.Split(strings.Split(parts[0], ",")[0], "=")[1]), utils.AtoiNoErr(strings.Split(parts[0], ",")[1])},
			vel: Coordinate{utils.AtoiNoErr(strings.Split(strings.Split(parts[1], ",")[0], "=")[1]), utils.AtoiNoErr(strings.Split(parts[1], ",")[1])},
		}
	}

	return robots
}

func simulateAndCalculate(robots []Robot) int {
	// simulate for 100 seconds
	for s := 0; s < 100; s++ {
		moveRobots(robots)
	}

	// count robots in each quadrant
	q1, q2, q3, q4 := 0, 0, 0, 0
	midX, midY := 101/2, 103/2

	for _, r := range robots {
		// skip robots exactly on the middle lines
		if r.pos.x == midX && r.pos.y == midY {
			continue
		}

		// count robots in each quad
		if r.pos.x < midX && r.pos.y < midY {
			q1++ // Top-left
		} else if r.pos.x > midX && r.pos.y < midY {
			q2++ // Top-right
		} else if r.pos.x < midX && r.pos.y > midY {
			q3++ // Bottom-left
		} else if r.pos.x > midX && r.pos.y > midY {
			q4++ // Bottom-right
		}
	}

	return q1 * q2 * q3 * q4
}

func moveRobots(robots []Robot) {
	for i := range robots {
		// Update position
		robots[i].pos.x += robots[i].vel.x
		robots[i].pos.y += robots[i].vel.y

		// Handle wrapping
		robots[i].pos.x = ((robots[i].pos.x % 101) + 101) % 101
		robots[i].pos.y = ((robots[i].pos.y % 103) + 103) % 103
	}
}

func findEasterEgg(robots []Robot) int {
	// Try each second until we find the pattern
	for t := 1; t < 11000; t++ {
		moveRobots(robots)

		// Check for pattern
		if hasLongHorizontalLine(robots) {
			return t
		}
	}
	return -1
}

func hasLongHorizontalLine(robots []Robot) bool {
	// Count robots in each row
	rowCounts := make(map[int]map[int]bool)
	for i := 0; i < 103; i++ {
		rowCounts[i] = make(map[int]bool)
	}

	// Record robot positions by row
	for _, robot := range robots {
		rowCounts[robot.pos.y][robot.pos.x] = true
	}

	// Look for a row with many robots and long consecutive sequences
	for y := 0; y < 103; y++ {
		if len(rowCounts[y]) > 30 {
			consecutive := 0
			for x := 0; x < 100; x++ {
				if rowCounts[y][x] && rowCounts[y][x+1] {
					consecutive++
					if consecutive > 25 {
						return true
					}
				} else {
					consecutive = 0
				}
			}
		}
	}

	return false
}
