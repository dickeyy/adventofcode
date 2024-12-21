package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-20/input.txt")
	utils.Output(day20(i, p))
}

type Pos struct {
	row, col int
}

type Grid struct {
	cells      [][]byte
	start, end Pos
	rows       int
	cols       int
}

func day20(input string, part int) int {
	grid := parseInput(input)

	// get base path distances for all reachable points
	distances := bfs(grid)

	// find possible skips that save time
	timeSkips := make(map[int]int)
	maxSkipDist := 2 // how far we can skip in a single cheat
	if part == 2 {
		maxSkipDist = 20
	}

	// for each position we can reach
	for pos, dist := range distances {
		// look at all positions within skip distance
		for dr := -maxSkipDist; dr <= maxSkipDist; dr++ {
			for dc := -maxSkipDist; dc <= maxSkipDist; dc++ {
				skipPos := Pos{pos.row + dr, pos.col + dc}

				// skip if out of bounds
				if !isInBounds(skipPos, grid) {
					continue
				}

				skipDist := utils.Abs(dr) + utils.Abs(dc)
				if skipDist == 0 || skipDist > maxSkipDist {
					continue
				}

				// if we can reach the skip destination normally
				if endDist, ok := distances[skipPos]; ok {
					// calculate the time saved
					timeSaved := endDist - dist - skipDist
					if timeSaved > 0 {
						timeSkips[timeSaved]++
					}
				}
			}
		}
	}

	// count the skips that save >= 100 time
	res := 0
	for saved, count := range timeSkips {
		if saved >= 100 {
			res += count
		}
	}

	return res
}

func parseInput(input string) (g Grid) {
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[0])
	cells := make([][]byte, rows)

	g.rows = rows
	g.cols = cols

	for r := range lines {
		cells[r] = make([]byte, cols)
		for c, ch := range lines[r] {
			cells[r][c] = byte(ch)
			if ch == 'S' {
				g.start = Pos{r, c}
			} else if ch == 'E' {
				g.end = Pos{r, c}
			}
		}
	}

	g.cells = cells
	return
}

func bfs(g Grid) map[Pos]int {
	distances := make(map[Pos]int)
	q := []Pos{g.start}
	distances[g.start] = 0

	moves := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		for _, m := range moves {
			n := Pos{c.row + m.row, c.col + m.col}
			if !isInBounds(n, g) || g.cells[n.row][n.col] == '#' {
				continue
			}

			if _, visited := distances[n]; !visited {
				distances[n] = distances[c] + 1
				q = append(q, n)
			}
		}
	}

	return distances
}

func isInBounds(p Pos, g Grid) bool {
	return p.row >= 0 && p.row < g.rows && p.col >= 0 && p.col < g.cols
}
