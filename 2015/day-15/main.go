package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.txt")
	utils.Output(day15(i, p))
}

const TSP_CAPACITY = 100 // each cookie has room for 100 teaspoons

// each ingredients properties are per 1 teaspoon
type Ingredient struct {
	Name       string
	Capacity   int // how well it helps the cookie absorb milk
	Durability int // how well it keeps the cookie intact when full of milk
	Flavor     int // how tasty it makes the cookie
	Texture    int // how it improves the feel of the cookie
	Calories   int // how many calories it adds to the cookie
}

type Recipe struct {
	Amounts []int
}

func day15(input string, part int) int {
	ingredients := make(map[int]Ingredient)
	out := 0 // max score

	for i, line := range strings.Split(input, "\n") {
		ingredients[i] = parseIngredient(line)
	}

	recipe := Recipe{
		Amounts: make([]int, len(ingredients)),
	}

	// Try all possible combinations
	for a := 0; a <= TSP_CAPACITY; a++ {
		recipe.Amounts[0] = a
		for b := 0; b <= TSP_CAPACITY-a; b++ {
			recipe.Amounts[1] = b
			for c := 0; c <= TSP_CAPACITY-a-b; c++ {
				recipe.Amounts[2] = c
				// The remaining amount must go to d
				recipe.Amounts[3] = TSP_CAPACITY - a - b - c

				// Calculate score
				score, valid := calcScore(recipe, ingredients, part == 2)
				if valid && score > out {
					out = score
				}
			}
		}
	}

	return out
}

func parseIngredient(line string) Ingredient {
	var ingredient Ingredient

	// split the line by the colon
	ingredient.Name = strings.Split(line, ": ")[0]
	line = strings.Split(line, ": ")[1]

	for _, c := range strings.Split(line, ", ") {
		p := strings.Split(c, " ")[0] // name of the property
		v := strings.Split(c, " ")[1] // value of the property

		switch p {
		case "capacity":
			ingredient.Capacity, _ = strconv.Atoi(v)
		case "durability":
			ingredient.Durability, _ = strconv.Atoi(v)
		case "flavor":
			ingredient.Flavor, _ = strconv.Atoi(v)
		case "texture":
			ingredient.Texture, _ = strconv.Atoi(v)
		case "calories":
			ingredient.Calories, _ = strconv.Atoi(v)
		}
	}

	return ingredient
}

func calcScore(recipe Recipe, ingredients map[int]Ingredient, checkCals bool) (int, bool) {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	calories := 0

	// Calculate each property total
	for i, amount := range recipe.Amounts {
		capacity += amount * ingredients[i].Capacity
		durability += amount * ingredients[i].Durability
		flavor += amount * ingredients[i].Flavor
		texture += amount * ingredients[i].Texture
		calories += amount * ingredients[i].Calories
	}

	// If any property is negative, the recipe is invalid
	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0, false
	}

	// For part 2, check calories
	if checkCals && calories != 500 {
		return 0, false
	}

	return capacity * durability * flavor * texture, true
}
