package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-24/input.txt")
	utils.Output(day24(i, p))
}

type GateInfo struct {
	operation int // 0=AND, 1=OR, 2=XOR
	inputs    []string
	output    string
}

func day24(input string, part int) string {
	wires, gates := parseInput(input)

	if part == 1 {
		// Evaluate gates until no more can be evaluated
		for len(gates) > 0 {
			for wireName, gate := range gates {
				if canEvalGate(gate, wires) {
					wires[wireName] = evaluateGate(gate, wires)
					delete(gates, wireName)
				}
			}
		}

		// Collect z-wires in order
		var zWires []string
		for wire := range wires {
			if strings.HasPrefix(wire, "z") {
				zWires = append(zWires, wire)
			}
		}
		sort.Strings(zWires)

		// Build binary number from most significant to least significant bit
		result := 0
		for i := len(zWires) - 1; i >= 0; i-- {
			result = (result << 1) | wires[zWires[i]]
		}

		return strconv.Itoa(result)
	}

	// Part 2
	var swapped []string
	var carry string

	// Convert gates map to slice of strings for easier searching
	var gateStrings []string
	for wireName, gate := range gates {
		gateStr := fmt.Sprintf("%s %s %s -> %s",
			gate.inputs[0],
			[]string{"AND", "OR", "XOR"}[gate.operation],
			gate.inputs[1],
			wireName)
		gateStrings = append(gateStrings, gateStr)
	}

	// Check each bit position for adder structure
	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var m1, n1, r1, z1, c1 string

		// Find half adder gates
		m1 = find(fmt.Sprintf("x%s", n), fmt.Sprintf("y%s", n), "XOR", gateStrings)
		n1 = find(fmt.Sprintf("x%s", n), fmt.Sprintf("y%s", n), "AND", gateStrings)

		if carry != "" {
			// Full adder logic
			r1 = find(carry, m1, "AND", gateStrings)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(carry, m1, "AND", gateStrings)
			}

			z1 = find(carry, m1, "XOR", gateStrings)

			// Check for misplaced z wires
			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}
			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}
			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			c1 = find(r1, n1, "OR", gateStrings)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if carry == "" {
			carry = n1
		} else {
			carry = c1
		}
	}

	sort.Strings(swapped)
	return strings.Join(swapped, ",")

}

func parseInput(input string) (map[string]int, map[string]GateInfo) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	wires := make(map[string]int)
	gates := make(map[string]GateInfo)

	// Parse initial wire values
	for _, line := range strings.Split(strings.TrimSpace(parts[0]), "\n") {
		parts := strings.Split(line, ": ")
		wires[parts[0]] = utils.AtoiNoErr(parts[1])
	}

	// Parse gates
	for _, line := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		inputs := strings.Split(parts[0], " ")

		var operation int
		var ins []string

		if len(inputs) == 3 {
			switch inputs[1] {
			case "AND":
				operation = 0
			case "OR":
				operation = 1
			case "XOR":
				operation = 2
			}
			ins = []string{inputs[0], inputs[2]}
		}

		gates[parts[1]] = GateInfo{
			operation: operation,
			inputs:    ins,
			output:    parts[1],
		}
	}

	return wires, gates
}

func evaluateGate(gate GateInfo, wires map[string]int) int {
	in1 := wires[gate.inputs[0]]
	in2 := wires[gate.inputs[1]]

	switch gate.operation {
	case 0: // AND
		return in1 & in2
	case 1: // OR
		return in1 | in2
	case 2: // XOR
		return in1 ^ in2
	}
	return 0
}

func canEvalGate(gate GateInfo, wires map[string]int) bool {
	_, hasIn1 := wires[gate.inputs[0]]
	_, hasIn2 := wires[gate.inputs[1]]
	return hasIn1 && hasIn2
}

func find(a, b, operator string, gates []string) string {
	for _, gate := range gates {
		if strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", a, operator, b)) ||
			strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", b, operator, a)) {
			parts := strings.Split(gate, " -> ")
			return parts[len(parts)-1]
		}
	}
	return ""
}
