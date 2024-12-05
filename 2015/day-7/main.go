package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-7/input.txt")
	utils.Output(day7(i, p))
}

// ---

type Instruction struct {
	Operation string
	Inputs    []string
	Output    string
}

type Wire struct {
	Value    uint16
	Assigned bool
}

var instructions []Instruction
var wires map[string]*Wire

func day7(input string, part int) int {
	// Parse instructions only once
	if len(instructions) == 0 {
		for _, inst := range strings.Split(input, "\n") {
			instructions = append(instructions, parseInstruction(inst))
		}
	}

	// Keep a copy of original instructions
	originalInstructions := make([]Instruction, len(instructions))
	copy(originalInstructions, instructions)

	// Reset wires
	wires = make(map[string]*Wire)

	// For part 2, we need to override B and run the simulation 2 times
	if part == 2 {
		// Run the simulation originally to get value of A
		runSimulation()
		overrideB := wires["a"].Value

		// Reset for second run
		wires = make(map[string]*Wire)
		wires["b"] = &Wire{Value: overrideB, Assigned: true}

		// Reset instructions, excluding the one setting 'b'
		instructions = []Instruction{}
		for _, inst := range originalInstructions {
			if inst.Output != "b" {
				instructions = append(instructions, inst)
			}
		}
	}

	// Run the simulation
	runSimulation()

	if wire, exists := wires["a"]; exists {
		return int(wire.Value)
	}
	return -1 // error
}

func runSimulation() {
	for len(instructions) > 0 {
		remainingInstructions := []Instruction{}

		for _, inst := range instructions {
			if inst.Output == "b" && wires["b"] != nil && wires["b"].Assigned {
				// Skip instructions trying to assign to 'b' if it's already set (for part 2)
				continue
			}

			if !executeInstruction(inst) {
				// Instruction couldn't be executed yet
				remainingInstructions = append(remainingInstructions, inst)
			}
		}

		if len(remainingInstructions) == len(instructions) {
			// No progress made, might be a circular dependency
			panic("Circular dependency detected")
		}

		instructions = remainingInstructions
	}
}

func parseInstruction(line string) Instruction {
	// Parse the line and return an Instruction
	// Handle different formats like "123 -> x", "x AND y -> z", "NOT x -> y"
	parts := strings.Split(line, " -> ")
	inst := Instruction{
		Output: strings.TrimSpace(parts[1]),
	}

	left := strings.Fields(parts[0]) // split the line into words, trimming whitespace

	switch len(left) {
	case 1:
		// case where the line is a direct assignment (x -> y)
		inst.Inputs = []string{left[0]}
		inst.Operation = "ASSIGN"
	case 2:
		// case where the line is a bitwise NOT operation (NOT x -> y)
		inst.Inputs = []string{left[1]}
		inst.Operation = left[0]
	case 3:
		// case where the line is a bitwise operation (x AND y -> z)
		inst.Inputs = []string{left[0], left[2]}
		inst.Operation = left[1]
	default:
		// error
		panic("Invalid instruction " + line)
	}

	return inst
}

func executeInstruction(inst Instruction) bool {
	// given an instruction, execute it and return true if successful
	switch inst.Operation {
	case "ASSIGN":
		if val, ok := getValue(inst.Inputs[0]); ok {
			wires[inst.Output] = &Wire{Value: val, Assigned: true}
			return true
		}
	case "NOT":
		if val, ok := getValue(inst.Inputs[0]); ok {
			wires[inst.Output] = &Wire{Value: ^val & 0xFFFF, Assigned: true}
			return true
		}
	default: // AND, OR, LSHIFT, RSHIFT
		if len(inst.Inputs) != 2 {
			panic("Invalid number of inputs for operation " + inst.Operation)
		}
		val1, ok1 := getValue(inst.Inputs[0])
		val2, ok2 := getValue(inst.Inputs[1])
		if ok1 && ok2 {
			result := bitwiseOperation(inst.Operation, val1, val2)
			wires[inst.Output] = &Wire{Value: result, Assigned: true}
			return true
		}
	}
	return false
}

// gets the value of a given input string. could be an existing wire or a number
// returns false if the input string is not a wire or number (error)
func getValue(input string) (uint16, bool) {
	if val, err := strconv.ParseUint(input, 10, 16); err == nil {
		return uint16(val), true
	}

	if wire, exists := wires[input]; exists && wire.Assigned {
		return wire.Value, true
	}
	return 0, false
}

func bitwiseOperation(op string, a, b uint16) uint16 {
	switch op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "LSHIFT":
		return a << b
	case "RSHIFT":
		return a >> b
	case "NOT":
		return ^a & 0xFFFF // Ensures the result is 16-bit
	default:
		panic("Invalid bitwise operation " + op)
	}
}
