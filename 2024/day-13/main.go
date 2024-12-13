package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-13/input.txt")
	utils.Output(day13(i, p))
}

type ClawMachine struct {
	btnA  Coordinate
	btnB  Coordinate
	prize Coordinate
}

type Coordinate struct {
	x, y int
}

func day13(input string, part int) int {
	machines := parseInput(input, part)
	totalTokens := 0

	for _, machine := range machines {
		totalTokens += calculateTokens(machine)
	}

	return totalTokens
}

func parseInput(input string, part int) map[int]*ClawMachine {
	machines := strings.Split(input, "\n\n")
	cms := make(map[int]*ClawMachine, len(machines))

	for i, machine := range machines {
		cms[i] = &ClawMachine{}

		for _, p := range strings.Split(machine, "\n") {
			t, v := strings.Split(p, ": ")[0], strings.Split(p, ": ")[1]
			xStr, yStr := strings.Split(v, ", ")[0], strings.Split(v, ", ")[1]

			switch t {
			case "Button A":
				x := utils.AtoiNoErr(strings.TrimPrefix(xStr, "X+"))
				y := utils.AtoiNoErr(strings.TrimPrefix(yStr, "Y+"))
				cms[i].btnA = Coordinate{x, y}

			case "Button B":
				x := utils.AtoiNoErr(strings.TrimPrefix(xStr, "X+"))
				y := utils.AtoiNoErr(strings.TrimPrefix(yStr, "Y+"))
				cms[i].btnB = Coordinate{x, y}

			case "Prize":
				x := utils.AtoiNoErr(strings.TrimPrefix(xStr, "X="))
				y := utils.AtoiNoErr(strings.TrimPrefix(yStr, "Y="))

				if part == 1 {
					cms[i].prize = Coordinate{x, y}
				} else {
					cms[i].prize = Coordinate{x + 10000000000000, y + 10000000000000}
				}
			}
		}
	}

	return cms
}

// this is a brute force solution for part 1, preserving it cuz idk
// func bruteCalculateTokens(machine *ClawMachine) int {
// 	// try all combinations of button presses up to 100 each
// 	for a := 0; a <= 100; a++ {
// 		for b := 0; b <= 100; b++ {
// 			// caclulate resulting X and Y positions after pressing A and B
// 			xPos := a*machine.btnA.x + b*machine.btnB.x
// 			yPos := a*machine.btnA.y + b*machine.btnB.y

// 			// if this combination reaches the prize coords
// 			if xPos == machine.prize.x && yPos == machine.prize.y {
// 				return (3 * a) + b // 3 tokens per A press and 1 token per B press
// 			}
// 		}
// 	}

// 	// if no combination of presses works, return 0
// 	return 0
// }

func calculateTokens(machine *ClawMachine) int {
	// use Cramer's rule to solve the system of equations
	a, b := solveEquation(machine)

	// if no valid solution, return 0
	if a <= 0 || b <= 0 {
		return 0
	}

	// calculate the total tokens (3 per A press, 1 per B press)
	return (3 * a) + b
}

func solveEquation(m *ClawMachine) (int, int) {
	// using cramer's rule:
	// for system of eqs:
	// ax*A + bx*B = px (x eq)
	// ay*A + by*B = py (y eq)

	// determinant of coefficients matrix
	d := m.btnA.x*m.btnB.y - m.btnA.y*m.btnB.x

	// determinants for A and B
	d1 := m.prize.x*m.btnB.y - m.prize.y*m.btnB.x
	d2 := m.prize.y*m.btnA.x - m.prize.x*m.btnA.y

	// check if we have whole number sols
	if d1%d != 0 || d2%d != 0 {
		return 0, 0
	}

	return d1 / d, d2 / d
}
