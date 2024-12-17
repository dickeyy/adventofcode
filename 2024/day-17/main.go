package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

/*
	This program simulates a simple 3-bit computer with three general-purpose registers (A, B, C).
	It executes a sequence of 3-bit instructions (0-7) where each instruction has an operand.
	The computer performs bitwise operations (XOR, right shifts) and basic arithmetic (modulo).

	Main operations:
		- Register manipulation through bitwise operations
		- 3-bit instruction decoding (opcode + operand pairs)
		- Dynamic operand resolution (direct values or register contents)
		- Output generation through modulo 8 operations

	Part 1: Executes the program and collects modulo 8 outputs
	Part 2: Uses binary search with 3-bit shifts to find the smallest initial A register value
    	    that makes the program output its own instructions

	Note: We could use some sort of Z3 solver for part 2, however, 1) idk how to do that,
		  and 2) brute force works just fine. :)
*/

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-17/input.txt")
	utils.Output(day17(i, p))
}

func day17(input string, part int) string {
	a, b, c, program := parseRegisters(input)

	if part == 1 {
		// part 1: run program and return comma-separated output
		return utils.IntArrayToString(runProgram(a, b, c, program), ",")
	}

	// part 2: find lowest value of register A that makes the program output itself
	a = 0 // the initial value of register A doesnt matter here, so we can reset it
	for pos := len(program) - 1; pos >= 0; pos-- {
		a <<= 3 // shift left by 3 bits
		for !slices.Equal(runProgram(a, b, c, program), program[pos:]) {
			a++
		}
	}

	return strconv.Itoa(a) // return a string since part 1 needs a string
}

func parseRegisters(input string) (int, int, int, []int) {
	a := 0
	b := 0
	c := 0
	program := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "Register A:") {
			a = utils.AtoiNoErr(strings.TrimSpace(line[12:]))
		} else if strings.HasPrefix(line, "Register B:") {
			b = utils.AtoiNoErr(strings.TrimSpace(line[12:]))
		} else if strings.HasPrefix(line, "Register C:") {
			c = utils.AtoiNoErr(strings.TrimSpace(line[12:]))
		} else if strings.HasPrefix(line, "Program:") {
			program = utils.GetIntsInString(strings.TrimSpace(line[9:]))
		}
	}

	return a, b, c, program
}

func runProgram(a, b, c int, program []int) []int {
	out := make([]int, 0)
	// for each isntruction pointer
	for ip := 0; ip < len(program); ip += 2 {
		opcode, operand := program[ip], program[ip+1]
		// Process combo operand
		value := operand
		switch operand {
		case 4:
			value = a
		case 5:
			value = b
		case 6:
			value = c
		}

		// Execute instruction
		switch opcode {
		case 0: // adv - divide A by 2^value
			a >>= value
		case 1: // bxl - XOR B with literal
			b ^= operand
		case 2: // bst - set B to value mod 8
			b = value % 8
		case 3: // jnz - jump if A is not zero
			if a != 0 {
				ip = operand - 2
			}
		case 4: // bxc - XOR B with C
			b ^= c
		case 5: // out - output value mod 8
			out = append(out, value%8)
		case 6: // bdv - divide A by 2^value, store in B
			b = a >> value
		case 7: // cdv - divide A by 2^value, store in C
			c = a >> value
		}
	}

	return out
}
