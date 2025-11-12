package utils

import "strings"

// grid represents a 2D grid of bytes
type Grid struct {
	Cells [][]byte
	Rows  int
	Cols  int
}

// point represents a coordinate in the grid
type Point struct {
	Row, Col int
}

// direction vectors for movement
var (
	Up        = Point{-1, 0}
	Down      = Point{1, 0}
	Left      = Point{0, -1}
	Right     = Point{0, 1}
	UpLeft    = Point{-1, -1}
	UpRight   = Point{-1, 1}
	DownLeft  = Point{1, -1}
	DownRight = Point{1, 1}
)

// CardinalDirections returns the 4 main directions (no diagonals)
func CardinalDirections() []Point {
	return []Point{Up, Down, Left, Right}
}

// AllDirections returns all 8 directions
func AllDirections() []Point {
	return []Point{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}
}

// ParseGrid creates a grid from a string input
func ParseGrid(input string) Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	if rows == 0 {
		return Grid{}
	}

	cols := len(lines[0])
	cells := make([][]byte, rows)

	for i, line := range lines {
		cells[i] = []byte(line)
		// ensure all rows have the same length
		if len(line) != cols {
			// pad shorter lines or truncate longer lines
			pad := make([]byte, cols-len(line))
			for j := range pad {
				pad[j] = ' '
			}
			cells[i] = append(cells[i], pad...)
		} else {
			cells[i] = cells[i][:cols]
		}
	}

	return Grid{
		Cells: cells,
		Rows:  rows,
		Cols:  cols,
	}
}

// IsInBounds checks if a point is within the grid boundaries
func (g *Grid) IsInBounds(p Point) bool {
	return p.Row >= 0 && p.Row < g.Rows && p.Col >= 0 && p.Col < g.Cols
}

// Get returns the byte at the given posiiton, or 0 if out of bounds
func (g *Grid) Get(p Point) byte {
	if !g.IsInBounds(p) {
		return 0
	}
	return g.Cells[p.Row][p.Col]
}

// Set sets the byte at the given position (no-op if out of bounds)
func (g *Grid) Set(p Point, val byte) {
	if g.IsInBounds(p) {
		g.Cells[p.Row][p.Col] = val
	}
}

// Add returns a new point that is the sum of p1 and p2
func (p Point) Add(other Point) Point {
	return Point{p.Row + other.Row, p.Col + other.Col}
}

// Neighbors returns all valid neighboring points in given directions
func (g *Grid) Neighbors(p Point, dirs []Point) []Point {
	var neighbors []Point
	for _, dir := range dirs {
		neighbor := p.Add(dir)
		if g.IsInBounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

// CardinalNeighbors Returns the 4 cardinal direction neighbors
func (g *Grid) CardinalNeighbors(p Point) []Point {
	return g.Neighbors(p, CardinalDirections())
}

// AllNeighbors returns all 8 direction neighbors
func (g *Grid) AllNeighbors(p Point) []Point {
	return g.Neighbors(p, AllDirections())
}

// FindAll returns all poisitions where the cell matches the given value
func (g *Grid) FindAll(target byte) []Point {
	var positions []Point
	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			if g.Cells[r][c] == target {
				positions = append(positions, Point{r, c})
			}
		}
	}
	return positions
}

// FindFirst returns the first position where the cell matches a vien value
func (g *Grid) FindFirst(target byte) (Point, bool) {
	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			if g.Cells[r][c] == target {
				return Point{r, c}, true
			}
		}
	}
	return Point{}, false
}

// String returns a string representation of the grid
func (g *Grid) String() string {
	var sb strings.Builder
	for r := 0; r < g.Rows; r++ {
		sb.Write(g.Cells[r])
		if r < g.Rows-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

// Clone returns a deep copy of the grid
func (g *Grid) Clone() Grid {
	newCells := make([][]byte, g.Rows)
	for r := 0; r < g.Rows; r++ {
		newCells[r] = make([]byte, g.Cols)
		copy(newCells[r], g.Cells[r])
	}
	return Grid{
		Cells: newCells,
		Rows:  g.Rows,
		Cols:  g.Cols,
	}
}

// ManhattanDistance returns the Manhattan distance between two points
func (p Point) ManhattanDistance(other Point) int {
	return Abs(p.Row-other.Row) + Abs(p.Col-other.Col)
}
