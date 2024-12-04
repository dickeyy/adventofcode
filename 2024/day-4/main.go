package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.txt")
	utils.Output(day4(i, p))
}

type Direction struct {
	row int
	col int
}

var dirs = []Direction{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // down
	{-1, 0},  // up
	{1, 1},   // down-right
	{1, -1},  // down-left
	{-1, 1},  // up-right
	{-1, -1}, // up-left
}

func day4(input string, part int) int {
	grid := createGrid(input)

	if part == 1 {
		return findPatternsPt1(grid)
	}

	return findPatternsPt2(grid)
}

func createGrid(input string) [][]rune {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}
	return grid
}

func isValid(row, col, rows, cols int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func checkXMASPattern(grid [][]rune, startRow, startCol int, dir Direction, rows, cols int) bool {
	// check if all 4 pos. will be within bounds
	for i := 0; i < 4; i++ {
		newRow := startRow + (dir.row * i)
		newCol := startCol + (dir.col * i)

		if !isValid(newRow, newCol, rows, cols) {
			return false
		}
	}

	pattern := "XMAS"
	for i := 0; i < len(pattern); i++ {
		currentRow := startRow + (dir.row * i)
		currentCol := startCol + (dir.col * i)

		if grid[currentRow][currentCol] != rune(pattern[i]) {
			return false
		}
	}

	return true
}

func findPatternsPt1(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	// for each pos. in the grid
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// only check poss. that start with "X"
			if grid[row][col] == 'X' {
				for _, dir := range dirs {
					if checkXMASPattern(grid, row, col, dir, rows, cols) {
						count++
					}
				}
			}
		}
	}

	return count
}

func findPatternsPt2(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	// skip edges since we need room for diagonals
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if grid[row][col] == 'A' {
				if checkXMASCross(grid, row, col, rows, cols) {
					count++
				}
			}
		}
	}

	return count
}

func checkXMASCross(grid [][]rune, centerRow, centerCol, rows, cols int) bool {
	// get diagonals
	ul := grid[centerRow-1][centerCol-1] // upper-left
	ur := grid[centerRow-1][centerCol+1] // upper-right
	ll := grid[centerRow+1][centerCol-1] // lower-left
	lr := grid[centerRow+1][centerCol+1] // lower-right

	// check if either diagonal forms MAS or SAM
	isValidDiag := func(first, last rune) bool {
		return (first == 'M' && last == 'S') || (first == 'S' && last == 'M')
	}

	return isValidDiag(ul, lr) && isValidDiag(ur, ll)
}
