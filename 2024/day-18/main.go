package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-18/input.txt")
	utils.Output(day18(i, p))
}

type Coord struct {
	x, y int
}

// priority queue types for Dijkstra's algorithm
type Item struct {
	coord    Coord
	priority int
	index    int
}

type PQueue []*Item

// direction vectors for movement (up, down, left, right)
var dirs = []Coord{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func day18(input string, part int) string {
	coords := parseInput(input)

	// crate a grid and mark corrupted spaces
	grid := make(map[Coord]bool)
	maxX, maxY := 70, 70

	if part == 1 {
		// only process first 1024 bytes
		numBytes := 1024
		if len(coords) > numBytes {
			coords = coords[:numBytes]
		}

		// mark corrupted spaces
		for _, coord := range coords {
			grid[coord] = true
		}

		// find the shortest path using Dijkstra's algorithm
		// return as a string because part 2 needs a string
		return strconv.Itoa(findShortestPath(grid, Coord{0, 0}, Coord{maxX, maxY}))
	}

	// for part 2, process bytes one at a time
	for _, coord := range coords {
		// add current byte to grid
		grid[coord] = true

		// check if path still exists
		if !pathExists(grid, Coord{0, 0}, Coord{maxX, maxY}) {
			// return the coordinates that blocked the path
			return fmt.Sprintf("%d,%d", coord.x, coord.y)
		}
	}

	return "erm what the sigma"
}

func parseInput(input string) []Coord {
	var coords []Coord
	for _, line := range strings.Split(input, "\n") {
		n := utils.GetIntsInString(line)
		coords = append(coords, Coord{n[0], n[1]})
	}
	return coords
}

// Dijkstra's algorithm
func findShortestPath(grid map[Coord]bool, start, end Coord) int {
	// init distances
	dist := make(map[Coord]int)
	dist[start] = 0

	// create a priority queue
	pq := make(PQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{coord: start, priority: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)

		// if we've reached the end, return the distance
		if curr.coord == end {
			return curr.priority
		}

		// if we've found a longer path to this point, skip
		if curr.priority > dist[curr.coord] {
			continue
		}

		// check all possible moves
		for _, dir := range dirs {
			next := Coord{curr.coord.x + dir.x, curr.coord.y + dir.y}

			// check if move is valid
			if isValidMove(next, grid, end) {
				newDist := curr.priority + 1
				// if we found a shorter path or havent't been here before
				if d, exists := dist[next]; !exists || newDist < d {
					dist[next] = newDist
					heap.Push(&pq, &Item{coord: next, priority: newDist})
				}
			}
		}
	}

	// if no path found
	return -1
}

// Part 2 use BFS to find if a path exists
func pathExists(grid map[Coord]bool, start, end Coord) bool {
	visited := make(map[Coord]bool)
	queue := []Coord{start}
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return true
		}

		for _, dir := range dirs {
			next := Coord{curr.x + dir.x, curr.y + dir.y}
			if isValidMove(next, grid, end) && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}

func isValidMove(pos Coord, grid map[Coord]bool, end Coord) bool {
	// check bounds
	if pos.x < 0 || pos.y < 0 || pos.x > end.x || pos.y > end.y {
		return false
	}

	// check if space is corrupted
	if grid[pos] {
		return false
	}

	return true
}

// --- Priority Queue Stuff (probably should move to a util but whatever for now) ---
func (pq PQueue) Len() int           { return len(pq) }
func (pq PQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
