package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	part := utils.GetPart()

	data := utils.ReadFile("./data.txt")
	utils.Output(day9(data, part))
}

// ---

type City string
type Distance int

type Graph struct {
	// Cities stores unique city names
	Cities []City
	// Distances stores the distance between any two cities
	// Key format: "CityA:CityB" where CityA < CityB lexicographically
	Distances map[string]Distance
}

func day9(data string, part int) int {
	graph := newGraph()
	graph.populate(data)

	return int(findPath(graph, part == 1))
}

func findPath(g *Graph, shortest bool) Distance {
	if len(g.Cities) == 0 {
		return 0
	}

	// track visted cities
	visited := make(map[City]bool)
	// Initialize outDist based on whether we're finding shortest or longest path
	var outDist Distance
	if shortest {
		outDist = Distance(math.MaxInt32)
	} else {
		outDist = Distance(math.MinInt32) // <-- This is the change needed
	}

	// try each city as a starting point
	for _, start := range g.Cities {
		visited[start] = true
		if shortest {
			dist := findShortestPathRecursive(g, start, visited, 1, 0)
			if dist < outDist {
				outDist = dist
			}
		} else {
			dist := findLongestPathRecursive(g, start, visited, 1, 0)
			if dist > outDist {
				outDist = dist
			}
		}
		visited[start] = false
	}

	return outDist
}
func findShortestPathRecursive(g *Graph, currentCity City, visited map[City]bool, citiesVisited int, currentDistance Distance) Distance {
	// if we've visted all cities we're done
	if citiesVisited == len(g.Cities) {
		return currentDistance
	}

	minDist := Distance(math.MaxInt32)

	// try each unvisted city as the next stop
	for _, nextCity := range g.Cities {
		if !visited[nextCity] {
			dist, exists := g.getDistance(currentCity, nextCity)
			if !exists {
				continue
			}

			visited[nextCity] = true
			totalDist := findShortestPathRecursive(g, nextCity, visited, citiesVisited+1, currentDistance+dist)
			if totalDist < minDist {
				minDist = totalDist
			}
			visited[nextCity] = false
		}
	}

	return minDist
}

func findLongestPathRecursive(g *Graph, currentCity City, visited map[City]bool, citiesVisited int, currentDistance Distance) Distance {
	// if we've visited all cities, we're done
	if citiesVisited == len(g.Cities) {
		return currentDistance
	}

	maxDist := Distance(math.MinInt32)

	// try each unvisted city as the next stop
	for _, nextCity := range g.Cities {
		if !visited[nextCity] {
			dist, exists := g.getDistance(currentCity, nextCity)
			if !exists {
				continue
			}

			visited[nextCity] = true
			totalDist := findLongestPathRecursive(g, nextCity, visited, citiesVisited+1, currentDistance+dist)
			if totalDist > maxDist {
				maxDist = totalDist
			}
			visited[nextCity] = false
		}
	}

	return maxDist
}

func newGraph() *Graph {
	return &Graph{
		Cities:    make([]City, 0),
		Distances: make(map[string]Distance),
	}
}

// populate takes in data and populates the graph
func (g *Graph) populate(data string) error {
	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " = ")
		dist, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}

		cities := strings.Split(parts[0], " to ")
		g.setDistance(City(cities[0]), City(cities[1]), Distance(dist))
	}

	return nil
}

// makeKey creates a consistent key for the distance map
// Always puts citites in lexicographical order to ensure consistent lookup
func makeKey(city1, city2 City) string {
	if city1 < city2 {
		return string(city1) + ":" + string(city2)
	}
	return string(city2) + ":" + string(city1)
}

// addCity adds a new city to the graph if it doesn't already exist
func (g *Graph) addCity(city City) {
	// check if city already exists
	for _, c := range g.Cities {
		if c == city {
			return
		}
	}
	g.Cities = append(g.Cities, city)
}

// setDistance sets the distance between two cities
func (g *Graph) setDistance(city1, city2 City, distance Distance) {
	g.addCity(city1)
	g.addCity(city2)
	key := makeKey(city1, city2)
	g.Distances[key] = distance
}

// getDistance gets the distance between any two cities
func (g *Graph) getDistance(City1, City2 City) (Distance, bool) {
	key := makeKey(City1, City2)
	distance, exists := g.Distances[key]
	return distance, exists
}
