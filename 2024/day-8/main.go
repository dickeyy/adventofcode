package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-8/input.txt")
	utils.Output(day8(i, p))
}

// for each pair, we need to find TWO antinodes
// (one on each side of the line)
//
// if antenna1 is at (x1, y1) and antenna2 is at (x2, y2)
// we need to find points P where:
// 1. P is collinear with antenna1 and antenna2
// 2. distance(P, antenna1) = 2 * distance(P, antenna2)
//    OR
//    distance(P, antenna2) = 2 * disntance(P, antenna1)

// this involves extending the line in both directions
// and finding the points that satisfy the 2:1 distance ratio

type Grid [][]rune
type Frequencies map[rune][]Position

type Position struct {
	x, y int
}

func day8(input string, part int) int {
	grid := parseGrid(input)
	frequencies := getFrequencies(grid)

	if part == 1 {
		return countAntinodes(grid, frequencies, false)
	}

	return countAntinodes(grid, frequencies, true)
}

// parseGrid converts the input string into a 2D rune slice
func parseGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line) // convert string directly into rune slice
	}
	return grid
}

// getFrequencies collects all antennas positions by their frequency and returns a map of locations
func getFrequencies(grid Grid) Frequencies {
	frequencies := make(Frequencies)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			char := grid[y][x]
			if char != '.' { // ignores filler spaces
				frequencies[char] = append(frequencies[char], Position{x, y})
			}
		}
	}

	return frequencies
}

func countAntinodes(grid Grid, frequencies map[rune][]Position, harmonic bool) int {
	antinodes := make(map[Position]bool) // track unique antinode positions

	// find antinodes for each frequency with 2+ antennas
	for _, pos := range frequencies {
		if len(pos) < 2 {
			continue
		}

		if harmonic {
			// Part 2: Check every point in the grid for collinearity
			for y := 0; y < len(grid); y++ {
				for x := 0; x < len(grid[y]); x++ {
					p := Position{x, y}
					// Check if this point is collinear with any pair of antennas
					for i := 0; i < len(pos); i++ {
						for j := i + 1; j < len(pos); j++ {
							if isCollinear(p, pos[i], pos[j]) {
								antinodes[p] = true
								break
							}
						}
					}
				}
			}
		} else {
			// Part 1: Original distance-based rules
			for i := 0; i < len(pos); i++ {
				for j := i + 1; j < len(pos); j++ {
					newAntinodes := findAntinodes(grid, pos[i], pos[j])
					for _, antinode := range newAntinodes {
						if isInBounds(antinode, grid) {
							antinodes[antinode] = true
						}
					}
				}
			}
		}
	}

	return len(antinodes)
}

// findAntinodes finds antinodes between any 2 antenna positions
func findAntinodes(grid Grid, a, b Position) []Position {
	antinodes := make([]Position, 0)

	// Check every point in the grid
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			p := Position{x, y}
			if isAntinode(p, a, b) {
				antinodes = append(antinodes, p)
			}
		}
	}

	return antinodes
}

// check if a point is a valid antinode for two antennas
func isAntinode(p, a, b Position) bool {
	if !isCollinear(p, a, b) {
		return false
	}

	// Calculate squared distances
	dap := distanceSquared(a, p)
	dbp := distanceSquared(b, p)

	// Check if either distance is twice the other
	// Note: We compare squares, so it's 4 times instead of 2 times
	return dap == 4*dbp || dbp == 4*dap
}

// helper to check if pos is within grid bounds
func isInBounds(pos Position, grid Grid) bool {
	return pos.y >= 0 && pos.y < len(grid) &&
		pos.x >= 0 && pos.x < len(grid[0])
}

// helper func to check if three points are collinear
func isCollinear(p, a, b Position) bool {
	// using the slop comparison method
	// (y2-y1)/(x2-x1) = (y3-y1)/(x3-x1)
	// cross multiply to avoid division by 0
	dx1 := a.x - p.x
	dy1 := a.y - p.y
	dx2 := b.x - p.x
	dy2 := b.y - p.y
	return dx1*dy2 == dx2*dy1
}

// helper func to calculate squared distance between two points
func distanceSquared(p1, p2 Position) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return dx*dx + dy*dy
}
