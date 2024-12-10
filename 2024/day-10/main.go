package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-10/input.txt")
	utils.Output(day10(i, p))
}

func day10(input string, part int) int {
	grid := parseInput(input)
	trailheads := findTrailHeads(grid)

	out := 0

	for _, trailhead := range trailheads {
		visited := make(map[[2]int]bool)
		if part == 1 {
			out += bfs(grid, trailhead)
		} else {
			out += countPaths(grid, trailhead, visited)
		}
	}

	return out
}

func parseInput(input string) [][]int {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(rows))

	for i, row := range rows {
		grid[i] = make([]int, len(row))
		for j, cell := range strings.Split(row, "") {
			grid[i][j] = utils.AtoiNoErr(cell)
		}
	}

	return grid
}

func findTrailHeads(grid [][]int) [][2]int {
	var trailheads [][2]int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}

	return trailheads
}

func bfs(grid [][]int, start [2]int) int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true
	score := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			ni, nj := curr[0]+d[0], curr[1]+d[1]
			if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) {
				next := [2]int{ni, nj}
				if !visited[next] && grid[ni][nj] == grid[curr[0]][curr[1]]+1 {
					visited[next] = true
					queue = append(queue, next)
					if grid[ni][nj] == 9 {
						score++
					}
				}
			}
		}
	}

	return score
}

func countPaths(grid [][]int, pos [2]int, visited map[[2]int]bool) int {
	if grid[pos[0]][pos[1]] == 9 {
		return 1 // reached the end of a valid path
	}

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited[pos] = true
	totalPaths := 0

	for _, d := range directions {
		ni, nj := pos[0]+d[0], pos[1]+d[1]
		next := [2]int{ni, nj}

		// check bounds and ensude path inrement is valid
		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) {
			if !visited[next] && grid[ni][nj] == grid[pos[0]][pos[1]]+1 {
				totalPaths += countPaths(grid, next, visited)
			}
		}
	}

	visited[pos] = false // backtrack
	return totalPaths
}
