package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./data.txt")
	utils.Output(pafh(data, part))
}

func pafh(data string, part int) int {
	grid := make([][]int, 1000)

	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, inst := range strings.Split(data, "\n") {
		op, x1, y1, x2, y2 := processInstruction(inst)

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if part == 1 {
					if op == "on" {
						grid[x][y] = 1
					} else if op == "off" {
						grid[x][y] = 0
					} else { // toggle
						grid[x][y] = 1 - grid[x][y]
					}
				} else { // part 2
					if op == "on" {
						grid[x][y]++
					} else if op == "off" {
						if grid[x][y] > 0 {
							grid[x][y]--
						}
					} else { // toggle
						grid[x][y] += 2
					}
				}
			}
		}
	}

	if part == 1 {
		return countOnLights(grid)
	}

	return sumBrightness(grid)
}

// helper func to process an instruction
func processInstruction(s string) (string, int, int, int, int) {
	// op can be "toggle" "turn off" or "turn on" but we only care about the second word (if there is one)
	var op string
	if strings.HasPrefix(s, "turn") {
		op = strings.Split(s, " ")[1]
	} else {
		op = strings.Split(s, " ")[0]
	}

	// remove the operation from the string
	if op == "toggle" {
		s = strings.TrimPrefix(s, "toggle ")
	} else {
		s = strings.TrimPrefix(s, "turn "+op+" ")
	}

	// coordinates of lights that will be affected
	x1, y1 := parseCoords(strings.Split(s, " ")[0])
	x2, y2 := parseCoords(strings.Split(s, " ")[2])

	return op, x1, y1, x2, y2
}

// helper func to parse coordinate strings into x and y ints
func parseCoords(s string) (int, int) {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return x, y
}

// helper func count the 1's in the grid
func countOnLights(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell > 0 {
				count++
			}
		}
	}
	return count
}

// helper func to sum the brightness of the grid
func sumBrightness(grid [][]int) int {
	sum := 0
	for _, row := range grid {
		for _, cell := range row {
			sum += cell
		}
	}
	return sum
}
