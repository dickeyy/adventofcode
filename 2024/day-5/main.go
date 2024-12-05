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

	i := utils.ReadFile("./input.txt")
	utils.Output(day5(i, p))
}

type Rule struct {
	before int
	after  int
}

type Graph map[int][]int

type Node struct {
	page     int
	inDegree int
	outNodes []int
}

func day5(input string, part int) int {
	out := 0

	rules, updates := parseInput(input)
	graph := buildGraph(rules)

	for _, update := range updates {
		if part == 1 {
			if isValidSequence(update, graph) {
				// get middle index, since we know sequences are guarenteed to have odd length
				out += update[len(update)/2]
			}
		} else {
			if !isValidSequence(update, graph) {
				// get the corrected ordering for invalid squences
				correctOrder := topologicalSort(update, graph)
				out += correctOrder[len(correctOrder)/2]
			}
		}
	}

	return out
}

func parseInput(input string) ([]Rule, [][]int) {
	var rules []Rule
	var updates [][]int

	parsingRules := true

	for _, line := range strings.Split(input, "\n") {
		// empty line separates the rules from the updates
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])
			rules = append(rules, Rule{before, after})
		} else {
			var update []int
			for _, num := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(num)
				update = append(update, n)
			}

			updates = append(updates, update)
		}
	}

	return rules, updates
}

func buildGraph(rules []Rule) Graph {
	graph := make(Graph)

	for _, rule := range rules {
		graph[rule.before] = append(graph[rule.before], rule.after)
	}

	return graph
}

// isValidSequence checks if a sequence satisifes all applicable rules
func isValidSequence(seq []int, graph Graph) bool {
	// create a map for quick lookups of positions
	positions := make(map[int]int)
	for i, page := range seq {
		positions[page] = i
	}

	// check each pair of pages in the sequence
	for i := 0; i < len(seq); i++ {
		page := seq[i]

		// if this page has rules
		if after, exists := graph[page]; exists {
			// check each page that should come after this one
			for _, mustBeAfter := range after {
				if pos, exists := positions[mustBeAfter]; exists {
					// check if it's actually after
					if pos <= i {
						return false
					}
				}
			}
		}
	}

	return true
}

// topologicalSort performs kahn's algorithm to sort pages
func topologicalSort(pages []int, fullGraph Graph) []int {
	// create subgraphs with only the pages we care about
	nodes := make(map[int]*Node)
	pageSet := make(map[int]bool)

	// init page set for quick lookups
	for _, page := range pages {
		pageSet[page] = true
		nodes[page] = &Node{
			page:     page,
			inDegree: 0,
			outNodes: []int{},
		}
	}

	// build the subgraphs and calculate in-degree
	for page := range pageSet {
		if outEdges, exists := fullGraph[page]; exists {
			for _, out := range outEdges {
				if pageSet[out] { // only include edges to pages in our set
					nodes[page].outNodes = append(nodes[page].outNodes, out)
					nodes[out].inDegree++
				}
			}
		}
	}

	// find starting nodes (in-degree = 0)
	var queue []int
	for page, node := range nodes {
		if node.inDegree == 0 {
			queue = append(queue, page)
		}
	}

	// sort queue to ensure deterministic ordering when multiple nodes have no dependencies
	sort.Slice(queue, func(i, j int) bool {
		return queue[i] > queue[j] // sort in descending order to match expected output
	})

	var result []int

	for len(queue) > 0 {
		// take first node from queue
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr)

		// update the neighbors
		node := nodes[curr]
		for _, next := range node.outNodes {
			nodes[next].inDegree--
			if nodes[next].inDegree == 0 {
				queue = append(queue, next)
				// re-sort queue after adding new node
				sort.Slice(queue, func(i, j int) bool {
					return queue[i] > queue[j]
				})
			}
		}
	}

	return result
}
