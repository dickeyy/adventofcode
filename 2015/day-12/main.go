package main

import (
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.json")
	utils.Output(day12(i, p))
}

func day12(input string, part int) int {
	out := 0
	if part == 1 {
		out = sumValues(input)
	} else {
		sum, err := parseAndSum(input)
		if err != nil {
			return 0
		}
		out = sum
	}
	return out
}

// part 1
func sumValues(data string) int {
	// Create regex pattern to match numbers (including negative ones)
	re := regexp.MustCompile(`-*\d+`)

	// Find all matches
	matches := re.FindAllString(data, -1)

	// Sum all numbers
	sum := 0
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			sum += num
		}
	}

	return sum
}

func sumWithoutRed(data interface{}) int {
	switch v := data.(type) {
	case float64: // JSON numbers are decoded as float64
		return int(v)

	case []interface{}: // Handle arrays
		sum := 0
		for _, item := range v {
			sum += sumWithoutRed(item)
		}
		return sum

	case map[string]interface{}: // Handle objects
		// Check if object has any "red" values
		for _, value := range v {
			if str, ok := value.(string); ok && str == "red" {
				return 0 // Ignore entire object
			}
		}

		// If no "red" found, sum all values
		sum := 0
		for _, value := range v {
			sum += sumWithoutRed(value)
		}
		return sum

	default:
		return 0 // Handle all other types (strings, etc.)
	}
}

func parseAndSum(input string) (int, error) {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return 0, err
	}
	return sumWithoutRed(data), nil
}
