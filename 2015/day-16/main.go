package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-16/input.txt")
	utils.Output(day16(i, p))
}

type Aunt struct {
	ID          int
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

func day16(input string, part int) int {
	// make a map of aunts
	aunts := make(map[int]Aunt)
	proximity := make(map[int]int) // associate id's with the number of categories they match with the desired aunt

	desiredAunt := Aunt{
		ID:          -1,
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	for _, line := range strings.Split(input, "\n") {
		a := parseAunt(line)
		aunts[a.ID] = a
	}

	for _, a := range aunts {
		// set the proximity of each aunt
		if part == 1 {
			proximity[a.ID] = mesureProximity(a, desiredAunt)
		} else {
			proximity[a.ID] = measureProximityRange(a, desiredAunt)
		}
	}

	// get the id of the aunt with the highest proximity
	var max int
	for id, proximity := range proximity {
		if proximity > max {
			max = proximity
			desiredAunt.ID = id
		}
	}

	return desiredAunt.ID
}

func parseAunt(line string) Aunt {
	a := Aunt{
		ID:          0, // always at least 1
		children:    -1,
		cats:        -1,
		samoyeds:    -1,
		pomeranians: -1,
		akitas:      -1,
		vizslas:     -1,
		goldfish:    -1,
		trees:       -1,
		cars:        -1,
		perfumes:    -1,
	}

	a.ID, _ = strconv.Atoi(strings.Split(strings.Split(line, ": ")[0], " ")[1])

	// trim the "Sue n: " from the beginning of each line
	line = strings.TrimPrefix(line, "Sue "+strconv.Itoa(a.ID)+": ")
	parts := strings.Split(line, ", ")

	for _, part := range parts[0:] {
		switch strings.Split(part, ": ")[0] {
		case "children":
			a.children, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "cats":
			a.cats, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "samoyeds":
			a.samoyeds, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "pomeranians":
			a.pomeranians, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "akitas":
			a.akitas, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "vizslas":
			a.vizslas, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "goldfish":
			a.goldfish, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "trees":
			a.trees, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "cars":
			a.cars, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		case "perfumes":
			a.perfumes, _ = strconv.Atoi(strings.Split(part, ": ")[1])
			continue
		default:
			continue
		}
	}

	return a
}

func mesureProximity(a Aunt, b Aunt) int {
	proximity := 0

	if a.children == b.children {
		proximity++
	}
	if a.cats == b.cats {
		proximity++
	}
	if a.samoyeds == b.samoyeds {
		proximity++
	}
	if a.pomeranians == b.pomeranians {
		proximity++
	}
	if a.akitas == b.akitas {
		proximity++
	}
	if a.vizslas == b.vizslas {
		proximity++
	}
	if a.goldfish == b.goldfish {
		proximity++
	}
	if a.trees == b.trees {
		proximity++
	}
	if a.cars == b.cars {
		proximity++
	}
	if a.perfumes == b.perfumes {
		proximity++
	}

	return proximity
}

func measureProximityRange(a Aunt, b Aunt) int {
	proximity := 0

	if a.children == b.children {
		proximity++
	}
	if a.cats > b.cats {
		proximity++
	}
	if a.samoyeds == b.samoyeds {
		proximity++
	}
	if a.pomeranians < b.pomeranians {
		proximity++
	}
	if a.akitas == b.akitas {
		proximity++
	}
	if a.vizslas == b.vizslas {
		proximity++
	}
	if a.goldfish < b.goldfish {
		proximity++
	}
	if a.trees > b.trees {
		proximity++
	}
	if a.cars == b.cars {
		proximity++
	}
	if a.perfumes == b.perfumes {
		proximity++
	}

	return proximity
}
