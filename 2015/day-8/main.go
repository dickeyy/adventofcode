package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	input := utils.ReadFile("./input.txt")
	utils.Output(day8(input, part))
}

func day8(input string, part int) int {
	out := 0

	for _, line := range strings.Split(input, "\n") {
		if part == 1 {
			out += len(line) - parseString(line)
		} else {
			out += len(encodeString(line)) - len(line)
		}
	}

	return out
}

func parseString(s string) int {
	/*
		For each string:
		- Code length is simple: len(string)
		- For memory length:
		  - Skip first and last quote
		  - When you see backslash:
		    - If followed by \ or ", count as one char
		    - If followed by x, skip next two chars and count as one char
		  - All other chars count as one
	*/

	mc := 0

	i := 1
	for i < len(s)-1 {
		if s[i] == '\\' { // backslash
			if i+1 >= len(s)-1 {
				break // safety check
			}

			switch s[i+1] {
			case '\\', '"':
				mc++
				i += 2 // skip both the backslash and the escaped char
			case 'x':
				if i+3 >= len(s)-1 {
					break // safety check for hex sequence
				}
				mc++
				i += 4 // skip \x and the two hex digits
			}
		} else {
			mc++
			i++
		}
	}

	return mc
}

func encodeString(s string) string {
	/*
		Start with opening quote
		For each character in original string:
		  - If it's a backslash: add another backslash
		  - If it's a quote: add backslash before it
		  - Otherwise: just add the character as-is
		Add closing quote
	*/

	ns := "\""

	for _, c := range s {
		if c == '\\' {
			ns += "\\\\"
		} else if c == '"' {
			ns += "\\\""
		} else {
			ns += string(c)
		}

	}

	ns += "\""
	return ns
}
