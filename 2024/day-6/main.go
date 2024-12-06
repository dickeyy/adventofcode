package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-6/input.txt")
	utils.Output(day6(i, p))
}

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

var (
	UP    = Direction{0, -1}
	DOWN  = Direction{0, 1}
	LEFT  = Direction{-1, 0}
	RIGHT = Direction{1, 0}
)

type Guard struct {
	pos Point
	dir Direction
}

type State struct {
	pos Point
	dir Direction
}

func day6(input string, part int) int {
	grid, guard := parseInput(input)

	if part == 1 {
		return simulateGuardPath(grid, guard, false)
	}

	return simLoop(grid, guard)
}

func (g *Guard) turnRight() {
	// turn right 90 degrees
	switch g.dir {
	case UP:
		g.dir = RIGHT
	case RIGHT:
		g.dir = DOWN
	case LEFT:
		g.dir = UP
	case DOWN:
		g.dir = LEFT
	}
}

func parseInput(input string) ([][]byte, Guard) {
	var grid [][]byte
	var guard Guard

	for y, line := range strings.Split(input, "\n") {
		row := make([]byte, len(line))
		for x, ch := range line {
			if ch == '^' {
				guard = Guard{Point{x, y}, UP}
				row[x] = '.' // replace the guard with an empty space
			} else {
				row[x] = byte(ch)
			}
		}
		grid = append(grid, row)
	}

	return grid, guard
}

// helper func to check if the position is in bounds
func isInBounds(p Point, grid [][]byte) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[0])
}

func simulateGuardPath(grid [][]byte, guard Guard, checkLoop bool) int {
	visited := make(map[Point]bool)
	visited[guard.pos] = true

	if checkLoop {
		states := make(map[State]struct{})
		states[State{guard.pos, guard.dir}] = struct{}{}

		for {
			next := Point{
				x: guard.pos.x + guard.dir.dx,
				y: guard.pos.y + guard.dir.dy,
			}

			if !isInBounds(next, grid) {
				return 0
			}

			if grid[next.y][next.x] == '#' {
				guard.turnRight()
			} else {
				guard.pos = next
			}

			state := State{guard.pos, guard.dir}
			if _, exists := states[state]; exists {
				return 1
			}
			states[state] = struct{}{}
		}
	}

	for {
		// calculate next position
		next := Point{
			x: guard.pos.x + guard.dir.dx,
			y: guard.pos.y + guard.dir.dy,
		}

		// check if guard would leave the area
		if !isInBounds(next, grid) {
			break
		}

		// check if there's an obstacle
		if grid[next.y][next.x] == '#' {
			guard.turnRight()
			continue
		}

		// move forward
		guard.pos = next
		visited[guard.pos] = true
	}

	return len(visited)
}

func simLoop(grid [][]byte, guard Guard) int {
	start := guard.pos
	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != '.' || (Point{x, y} == start) {
				continue
			}

			// create a copy of the grid with the new obstruction
			newGrid := make([][]byte, len(grid))
			for i := range grid {
				newGrid[i] = make([]byte, len(grid[i]))
				copy(newGrid[i], grid[i])
			}
			newGrid[y][x] = '#'

			if simulateGuardPath(newGrid, guard, true) > 0 {
				count++
			}
		}
	}
	return count
}
