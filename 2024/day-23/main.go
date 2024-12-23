package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-23/input.txt")
	utils.Output(day23(i, p))
}

type AdjList map[string]map[string]bool

func day23(input string, part int) string {
	adj := parseInput(input)
	if part == 1 {
		return strconv.Itoa(findTriangles(adj)) // return as string since part 2 returns a string
	}
	return strings.Join(findLargestConnected(adj), ",")
}

func parseInput(input string) AdjList {
	// build adjacency list
	adj := make(AdjList)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		nodes := strings.Split(line, "-")
		a, b := nodes[0], nodes[1]

		// init maps if needed
		if adj[a] == nil {
			adj[a] = make(map[string]bool)
		}
		if adj[b] == nil {
			adj[b] = make(map[string]bool)
		}

		// add bidirectional edge
		adj[a][b] = true
		adj[b][a] = true
	}

	return adj
}

func findTriangles(adj AdjList) int {
	// find triangles with at least one 't' node
	seen := make(map[string]bool)
	count := 0

	// for each node
	for node := range adj {
		// and each neighbor
		for neighbor := range adj[node] {
			if seen[neighbor+node] {
				continue
			}
			seen[node+neighbor] = true

			// for each neighbor of neighbor
			for non := range adj[neighbor] {
				if non == node {
					continue
				}

				// if forms a triangle and has 't' node
				if adj[non][node] {
					if strings.HasPrefix(node, "t") ||
						strings.HasPrefix(neighbor, "t") ||
						strings.HasPrefix(non, "t") {
						count++
					}
				}
			}
		}
	}

	return count / 3 // each triangle is 3 nodes
}

// findLargestConnected is just the Bron–Kerbosch algorithm
func findLargestConnected(adj AdjList) []string {
	// convert AdjList to the format needed for Bron–Kerbosch
	graph := make(map[string][]string)
	for node, neighbors := range adj {
		nodeNeighbors := make([]string, 0, len(neighbors))
		for neighbor := range neighbors {
			nodeNeighbors = append(nodeNeighbors, neighbor)
		}
		graph[node] = nodeNeighbors
	}

	// init sets for Bron-Kerbosch
	R := make(map[string]bool) // current clique being built
	P := make(map[string]bool) // prospective nodes
	X := make(map[string]bool) // excluded nodes

	// start with all vertices in P
	for v := range graph {
		P[v] = true
	}

	var maxClique map[string]bool

	// recursive Bron-Kerbosch implementation
	var bronKerbosch func(R, P, X map[string]bool)
	bronKerbosch = func(R, P, X map[string]bool) {
		if len(P) == 0 && len(X) == 0 {
			// found a maximal clique
			if len(R) > len(maxClique) {
				maxClique = make(map[string]bool)
				for v := range R {
					maxClique[v] = true
				}
			}
			return
		}

		for v := range P {
			// create new sets for recursive call
			newR := copySet(R)
			newR[v] = true

			// create new P and X sets containing only neighbors of v
			newP := make(map[string]bool)
			newX := make(map[string]bool)
			for _, n := range graph[v] {
				if P[n] {
					newP[n] = true
				}
				if X[n] {
					newX[n] = true
				}
			}

			bronKerbosch(newR, newP, newX)

			// move V from P to X
			delete(P, v)
			X[v] = true
		}
	}

	bronKerbosch(R, P, X)

	// convert maximal clique to slice
	res := make([]string, 0, len(maxClique))
	for v := range maxClique {
		res = append(res, v)
	}
	sort.Strings(res)
	return res
}

func copySet(s map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k, v := range s {
		res[k] = v
	}
	return res
}
