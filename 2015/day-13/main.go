package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2015/day-13/input.txt")
	utils.Output(day13(i, p))
}

// ---

// nested map of people and their happiness when seated next to another person
type Person string
type Happiness map[Person]map[Person]int

func day13(input string, part int) int {
	happiness := make(Happiness)
	people := map[Person]struct{}{}

	for _, line := range strings.Split(input, "\n") {
		p1, p2, num := parseLine(line)

		// add to happiness map
		people[p1] = struct{}{}
		people[p2] = struct{}{}

		// init happiness map
		if happiness[p1] == nil {
			happiness[p1] = make(map[Person]int)
		}
		happiness[p1][p2] = num
	}

	// Convert map to slice for permutations
	var peopleSlice []Person
	for person := range people {
		peopleSlice = append(peopleSlice, person)
	}

	// handle part 2
	if part == 2 {
		happiness, peopleSlice = addMe(happiness, peopleSlice)
	}

	// generate permuations
	arrangements := permute(peopleSlice)

	return calcMax(happiness, arrangements)
}

func parseLine(line string) (Person, Person, int) {
	// person 1 is the first word
	// person 2 is the last word (minus the period)
	p1, p2 := line[:strings.Index(line, " ")], line[strings.LastIndex(line, " ")+1:]
	// remove the period
	p2 = p2[:len(p2)-1]

	// get the number of happiness units change
	re := regexp.MustCompile(`-*\d+`) // get all numbers
	num, _ := strconv.Atoi(re.FindAllString(line, -1)[0])

	// get the direction of the change (gain or loss) and reflect it in the number
	if string(line[strings.Index(line, "would")+len("would") : strings.Index(line, "by")][1]) == "l" {
		num *= -1
	}

	return Person(p1), Person(p2), num
}

func permute(people []Person) [][]Person {
	if len(people) == 1 {
		return [][]Person{{people[0]}}
	}

	var result [][]Person
	first := people[0]
	rest := people[1:]

	// Get permutations of everything except first
	subPerms := permute(rest)

	// For each sub-permutation
	for _, perm := range subPerms {
		// Insert first at every possible position
		for i := 0; i <= len(perm); i++ {
			newPerm := make([]Person, len(perm)+1)
			copy(newPerm[:i], perm[:i])
			newPerm[i] = first
			copy(newPerm[i+1:], perm[i:])
			result = append(result, newPerm)
		}
	}
	return result
}

func calcMax(h Happiness, ars [][]Person) int {
	max := math.MinInt
	for _, a := range ars {
		total := 0

		for i := 0; i < len(a); i++ {
			curr := a[i]
			l := a[(i-1+len(a))%len(a)]
			r := a[(i+1)%len(a)]

			// add happiness from both neighbors
			total += h[curr][l]
			total += h[curr][r]
		}

		if total > max {
			max = total
		}
	}

	return max
}

// part 2
func addMe(h Happiness, ps []Person) (Happiness, []Person) {
	me := Person("me")
	// Add me to the slice
	nps := append(ps, me)

	// Initialize my happiness map
	h[me] = make(map[Person]int)

	// Add 0 happiness for all relationships involving me
	for _, p := range nps {
		if p != me {
			// me -> others
			h[me][p] = 0

			// others -> me
			if h[p] == nil {
				h[p] = make(map[Person]int)
			}
			h[p][me] = 0
		}
	}

	return h, nps
}
