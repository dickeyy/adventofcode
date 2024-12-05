package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-18/input.txt")
	utils.Output(day18(i, p))
}

const (
	gridSize = 100
	steps    = 100
)

type Grid [][]bool

func day18(input string, part int) int {
	g := newGrid()

	for row, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		for col, char := range line {
			g[row][col] = char == '#'
		}
	}

	if part == 2 {
		g[0][0] = true
		g[0][gridSize-1] = true
		g[gridSize-1][0] = true
		g[gridSize-1][gridSize-1] = true
	}

	for i := 0; i < steps; i++ {
		g = g.step(part)
	}

	return g.countLights()
}

func newGrid() Grid {
	grid := make([][]bool, gridSize)
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}
	return grid
}

func (g Grid) countNeighbors(row, col int) int {
	count := 0
	// chck all 8 adjacent positions
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue // skip the current position
			}
			nr, nc := row+dr, col+dc
			if nr >= 0 && nr < gridSize && nc >= 0 && nc < gridSize {
				if g[nr][nc] {
					count++
				}
			}
		}
	}
	return count
}

// step performs one animation step
func (g Grid) step(part int) Grid {
	ng := newGrid()

	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			neighbors := g.countNeighbors(row, col)

			if g[row][col] {
				// light is currently ojn
				ng[row][col] = neighbors == 2 || neighbors == 3
			} else {
				// light is currently off
				ng[row][col] = neighbors == 3
			}
		}
	}

	if part == 2 {
		ng[0][0] = true
		ng[0][gridSize-1] = true
		ng[gridSize-1][0] = true
		ng[gridSize-1][gridSize-1] = true
	}

	return ng
}

func (g Grid) countLights() int {
	count := 0
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if g[row][col] {
				count++
			}
		}
	}
	return count
}
