package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-15/input.txt")
	utils.Output(day15(i, p))
}

type Pair struct {
	r, c int
}

type Direction struct {
	dx, dy int
}

var directions = map[byte]Direction{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

// Part 1 structures
type Warehouse struct {
	moveSeq string
	boxes   map[Pair]struct{}
	robot   Pair
	walls   map[Pair]struct{}
	width   int
	height  int
}

// Part 2 structures
type BigBox struct {
	left  Pair
	right Pair
}

type BigWarehouse struct {
	moveSeq  string
	boxes    map[BigBox]struct{}
	boxParts map[Pair]BigBox
	robot    Pair
	walls    map[Pair]struct{}
	width    int
	height   int
}

func day15(input string, part int) int {
	if part == 1 {
		warehouse := parseWarehouse(input)
		return solveP1(warehouse)
	}
	warehouse := parseWarehouse(input)
	bigWarehouse := newBigWarehouse(warehouse)
	return solveP2(bigWarehouse)
}

func parseWarehouse(input string) Warehouse {
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")
	moves := strings.ReplaceAll(parts[1], "\n", "")

	w := Warehouse{
		moveSeq: moves,
		boxes:   make(map[Pair]struct{}),
		walls:   make(map[Pair]struct{}),
	}

	w.height = len(lines)
	w.width = len(lines[0])

	for r, line := range lines {
		for c, ch := range line {
			p := Pair{r, c}
			switch ch {
			case '#':
				w.walls[p] = struct{}{}
			case 'O':
				w.boxes[p] = struct{}{}
			case '@':
				w.robot = p
			}
		}
	}

	return w
}

func newBigWarehouse(w Warehouse) BigWarehouse {
	bw := BigWarehouse{
		moveSeq:  w.moveSeq,
		boxes:    make(map[BigBox]struct{}),
		boxParts: make(map[Pair]BigBox),
		walls:    make(map[Pair]struct{}),
		height:   w.height,
		width:    w.width * 2,
	}

	// Convert robot position
	bw.robot = Pair{w.robot.r, w.robot.c * 2}

	// Convert walls
	for wall := range w.walls {
		wall1 := Pair{wall.r, wall.c * 2}
		wall2 := Pair{wall1.r, wall1.c + 1}
		bw.walls[wall1] = struct{}{}
		bw.walls[wall2] = struct{}{}
	}

	// Convert boxes
	for box := range w.boxes {
		left := Pair{box.r, box.c * 2}
		right := Pair{box.r, left.c + 1}
		bigBox := BigBox{left, right}
		bw.boxParts[left] = bigBox
		bw.boxParts[right] = bigBox
		bw.boxes[bigBox] = struct{}{}
	}

	return bw
}

func getNextPair(p Pair, dir byte) Pair {
	d := directions[dir]
	return Pair{p.r + d.dy, p.c + d.dx}
}

// Part 1 functions
func canBoxMove(w *Warehouse, box Pair, dir byte) bool {
	next := getNextPair(box, dir)
	_, isWall := w.walls[next]
	if isWall {
		return false
	}
	_, isBox := w.boxes[next]
	if isBox {
		return canBoxMove(w, next, dir)
	}
	return true
}

func moveBoxes(w *Warehouse, box Pair, dir byte) {
	next := getNextPair(box, dir)
	_, isBox := w.boxes[next]
	if isBox {
		moveBoxes(w, next, dir)
	}
	delete(w.boxes, box)
	w.boxes[next] = struct{}{}
}

func move(w *Warehouse, dir byte) {
	next := getNextPair(w.robot, dir)
	_, isWall := w.walls[next]
	if isWall {
		return
	}
	_, isBox := w.boxes[next]
	if isBox && canBoxMove(w, next, dir) {
		moveBoxes(w, next, dir)
		w.robot = next
	} else if !isBox {
		w.robot = next
	}
}

// Part 2 functions
func canBigBoxMove(w *BigWarehouse, side Pair, dir byte) bool {
	bb := w.boxParts[side]
	left, right := bb.left, bb.right
	leftNext := getNextPair(left, dir)
	rightNext := getNextPair(right, dir)

	_, lWall := w.walls[leftNext]
	_, rWall := w.walls[rightNext]
	if lWall || rWall {
		return false
	}

	if dir == '<' {
		_, lBox := w.boxParts[leftNext]
		if lBox {
			return canBigBoxMove(w, leftNext, dir)
		}
		return true
	}

	if dir == '>' {
		_, rBox := w.boxParts[rightNext]
		if rBox {
			return canBigBoxMove(w, rightNext, dir)
		}
		return true
	}

	bbL, lBox := w.boxParts[leftNext]
	bbR, rBox := w.boxParts[rightNext]

	canMove := true
	if lBox {
		canMove = canMove && canBigBoxMove(w, leftNext, dir)
	}
	if rBox && bbL != bbR {
		canMove = canMove && canBigBoxMove(w, rightNext, dir)
	}
	return canMove
}

func bigBoxMove(w *BigWarehouse, side Pair, dir byte) {
	bb := w.boxParts[side]
	left, right := bb.left, bb.right
	leftNext, rightNext := getNextPair(left, dir), getNextPair(right, dir)

	if dir == '<' {
		_, lBox := w.boxParts[leftNext]
		if lBox {
			bigBoxMove(w, leftNext, dir)
		}
		delete(w.boxes, bb)
		delete(w.boxParts, left)
		delete(w.boxParts, right)
		bb.right = bb.left
		bb.left = leftNext
		w.boxes[bb] = struct{}{}
		w.boxParts[left] = bb
		w.boxParts[leftNext] = bb
		return
	}

	if dir == '>' {
		_, rBox := w.boxParts[rightNext]
		if rBox {
			bigBoxMove(w, rightNext, dir)
		}
		delete(w.boxes, bb)
		delete(w.boxParts, left)
		delete(w.boxParts, right)
		bb.left = bb.right
		bb.right = rightNext
		w.boxes[bb] = struct{}{}
		w.boxParts[right] = bb
		w.boxParts[rightNext] = bb
		return
	}

	bbL, lBox := w.boxParts[leftNext]
	bbR, rBox := w.boxParts[rightNext]

	if lBox {
		bigBoxMove(w, leftNext, dir)
	}
	if rBox && bbL != bbR {
		bigBoxMove(w, rightNext, dir)
	}

	delete(w.boxes, bb)
	delete(w.boxParts, left)
	delete(w.boxParts, right)
	bb.left = leftNext
	bb.right = rightNext
	w.boxes[bb] = struct{}{}
	w.boxParts[leftNext] = bb
	w.boxParts[rightNext] = bb
}

func bigMove(w *BigWarehouse, dir byte) {
	next := getNextPair(w.robot, dir)
	_, isWall := w.walls[next]
	if isWall {
		return
	}
	_, isBox := w.boxParts[next]
	if isBox && canBigBoxMove(w, next, dir) {
		bigBoxMove(w, next, dir)
		w.robot = next
	} else if !isBox {
		w.robot = next
	}
}

func solveP1(w Warehouse) int {
	moves := []byte(w.moveSeq)
	for _, moveDir := range moves {
		move(&w, moveDir)
	}

	score := 0
	for box := range w.boxes {
		score += 100*box.r + box.c
	}
	return score
}

func solveP2(w BigWarehouse) int {
	moves := []byte(w.moveSeq)
	for _, move := range moves {
		bigMove(&w, move)
	}

	score := 0
	for box := range w.boxes {
		score += 100*box.left.r + box.left.c
	}
	return score
}
