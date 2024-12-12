package main

import (
	"sort"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-12/input.txt")
	utils.Output(day12(i, p))
}

// Direction represents a cardinal direction in the grid
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// Coord represents a point in the grid
type Coord struct {
	row, col int
}

// Garden represents a connected group of plants
type Garden struct {
	symbol     rune
	size       int
	coords     map[Coord]bool
	boundaries map[Direction][]boundary
}

// boundary represents an edge of a garden
type boundary struct {
	row, col int
	counted  bool
}

func day12(input string, part int) int {
	grid := parseGrid(input)
	gardens := findGardens(grid)
	return calculateScore(gardens, part == 2)
}

func parseGrid(input string) [][]rune {
	ls := strings.Split(input, "\n")
	g := make([][]rune, len(ls))
	for i, l := range ls {
		g[i] = []rune(l)
	}
	return g
}

// find all gardens in the grid
func findGardens(grid [][]rune) []Garden {
	s := make(map[Coord]bool)
	var gs []Garden

	for row := range grid {
		for col := range grid[row] {
			coord := Coord{row, col}
			if !s[coord] {
				if g := exploreGarden(grid, coord, s); g.size > 0 {
					gs = append(gs, g)
				}
			}
		}
	}

	return gs
}

// flood fill to find connected regions
func exploreGarden(grid [][]rune, start Coord, seen map[Coord]bool) Garden {
	if seen[start] || !isValidCoord(grid, start) {
		return Garden{}
	}

	symbol := grid[start.row][start.col]
	garden := Garden{
		symbol:     symbol,
		coords:     make(map[Coord]bool),
		boundaries: make(map[Direction][]boundary),
	}

	// Initialize the queue with the starting point
	queue := []Coord{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if seen[curr] || !isValidCoord(grid, curr) || grid[curr.row][curr.col] != symbol {
			continue
		}

		seen[curr] = true
		garden.coords[curr] = true
		garden.size++

		// Check boundaries in all directions
		checkBoundary(grid, curr, symbol, North, &garden)
		checkBoundary(grid, curr, symbol, South, &garden)
		checkBoundary(grid, curr, symbol, East, &garden)
		checkBoundary(grid, curr, symbol, West, &garden)

		// Add adjacent cells to queue
		for _, next := range getAdjacent(curr) {
			if isValidCoord(grid, next) && grid[next.row][next.col] == symbol {
				queue = append(queue, next)
			}
		}
	}

	return garden
}

// check if a boundary is an edge
func checkBoundary(grid [][]rune, pos Coord, symbol rune, dir Direction, garden *Garden) {
	var isEdge bool
	switch dir {
	case North:
		isEdge = pos.row == 0 || grid[pos.row-1][pos.col] != symbol
	case South:
		isEdge = pos.row == len(grid)-1 || grid[pos.row+1][pos.col] != symbol
	case East:
		isEdge = pos.col == len(grid[0])-1 || grid[pos.row][pos.col+1] != symbol
	case West:
		isEdge = pos.col == 0 || grid[pos.row][pos.col-1] != symbol
	}

	if isEdge {
		garden.boundaries[dir] = append(garden.boundaries[dir], boundary{pos.row, pos.col, true})
	}
}

func calculateScore(gardens []Garden, isPart2 bool) int {
	total := 0
	for _, garden := range gardens {
		if isPart2 {
			pruneRedundantBoundaries(&garden)
		}
		boundaryCount := countValidBoundaries(garden)
		total += garden.size * boundaryCount
	}
	return total
}

func pruneRedundantBoundaries(garden *Garden) {
	for dir := range garden.boundaries {
		if dir == North || dir == South {
			pruneBoundariesAlongAxis(garden.boundaries[dir], true)
		} else {
			pruneBoundariesAlongAxis(garden.boundaries[dir], false)
		}
	}
}

func pruneBoundariesAlongAxis(bounds []boundary, sortByCol bool) {
	if sortByCol {
		sort.Slice(bounds, func(i, j int) bool { return bounds[i].col < bounds[j].col })
	} else {
		sort.Slice(bounds, func(i, j int) bool { return bounds[i].row < bounds[j].row })
	}

	for i := range bounds {
		var next int
		if sortByCol {
			next = bounds[i].col
		} else {
			next = bounds[i].row
		}

		for {
			next++
			found := false
			for j := range bounds {
				var matches bool
				if sortByCol {
					matches = bounds[j].col == next && bounds[j].row == bounds[i].row
				} else {
					matches = bounds[j].row == next && bounds[j].col == bounds[i].col
				}
				if matches {
					bounds[j].counted = false
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
	}
}

func countValidBoundaries(garden Garden) int {
	count := 0
	for _, bounds := range garden.boundaries {
		for _, b := range bounds {
			if b.counted {
				count++
			}
		}
	}
	return count
}

func isValidCoord(grid [][]rune, c Coord) bool {
	return c.row >= 0 && c.row < len(grid) && c.col >= 0 && c.col < len(grid[0])
}

func getAdjacent(c Coord) []Coord {
	return []Coord{
		{c.row - 1, c.col}, // North
		{c.row + 1, c.col}, // South
		{c.row, c.col + 1}, // East
		{c.row, c.col - 1}, // West
	}
}
