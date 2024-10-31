package main

import (
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.txt")
	utils.Output(day14(i, p))
}

const RACE_LENGTH = 2503

type Reindeer struct {
	Name       string // name. lol
	Speed      int    // the speed at which it can fly
	FlightTime int    // how long it can fly at its speed before rest
	Rest       int    // how long it must rest before it can fly again

	// New fields for part 2
	Distance    int // current distance
	Points      int // points accumulated
	TimeInCycle int // track where in cycle the reindeer is
}

func day14(input string, part int) int {
	var deers []Reindeer

	for _, line := range strings.Split(input, "\n") {
		deers = append(deers, parseReindeer(line))
	}

	if part == 1 {
		return raceDeer(deers)
	}

	// sim each second
	for second := 0; second < RACE_LENGTH; second++ {
		// update positions
		updatePositions(deers)
		// award points to leader(s)
		awardPoints(deers)
	}

	return findMaxPoints(deers)
}

func raceDeer(deers []Reindeer) int {
	furthestDistance := 0

	for _, deer := range deers {
		dist := calcDistance(deer)
		if dist > furthestDistance {
			furthestDistance = dist
		}
	}

	return furthestDistance
}

func parseReindeer(line string) Reindeer {
	r := Reindeer{}
	r.Name = line[:strings.Index(line, " can fly ")]
	nums := utils.GetIntsInString(line)
	r.Speed = nums[0]
	r.FlightTime = nums[1]
	r.Rest = nums[2]
	return r
}

func calcDistance(deer Reindeer) int {
	// First calculate complete cycles
	completeCycles := RACE_LENGTH / (deer.FlightTime + deer.Rest)
	distanceInCompleteCycles := completeCycles * deer.Speed * deer.FlightTime

	// Handle the remaining time
	remainingTime := RACE_LENGTH % (deer.FlightTime + deer.Rest)

	// For remaining time, we can only count flying time
	// and only up to FlightTime duration
	extraFlyingTime := min(remainingTime, deer.FlightTime)
	extraDistance := extraFlyingTime * deer.Speed

	return distanceInCompleteCycles + extraDistance
}

func updatePositions(deers []Reindeer) {
	for i := range deers {
		// calculate if deer is in flying period
		cycleLength := deers[i].FlightTime + deers[i].Rest
		timeInCycle := deers[i].TimeInCycle % cycleLength

		// if in flying period, add distance
		if timeInCycle < deers[i].FlightTime {
			deers[i].Distance += deers[i].Speed
		}

		deers[i].TimeInCycle++
	}
}

func awardPoints(deers []Reindeer) {
	maxDist := 0

	for _, d := range deers {
		if d.Distance > maxDist {
			maxDist = d.Distance
		}
	}

	for i := range deers {
		if deers[i].Distance == maxDist {
			deers[i].Points++
		}
	}
}

func findMaxPoints(deers []Reindeer) int {
	var max int
	for _, d := range deers {
		if d.Points > max {
			max = d.Points
		}
	}
	return max
}
