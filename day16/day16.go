package main

import (
	common "aoc2024/aoccommon"
	"container/heap"
	"fmt"
)

type PartialPath struct {
	pqIndex int
	cost int
	coord [2]int
	dxDy [2]int
	predecessors []*PartialPath
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*PartialPath

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].pqIndex = i
	pq[j].pqIndex = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PartialPath)
	item.pqIndex = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.pqIndex = -1 // for safety
	*pq = old[0 : n-1]
	return item
}


func printCostMap(partialPathsByPosition map[[2][2]int]*PartialPath, lines []string) {
	maxCost := 0
	for _, path := range partialPathsByPosition {
		if path.cost > maxCost {
			maxCost = path.cost
		}
	}
	cellWidth := len(string(maxCost)) + 3
	for y := range len(lines) {
		if y > 40 {
			continue
		}
		for x := range len(lines[0]) {
			if x < len(lines[0]) - 25 {
				continue
			}
			minCost := maxCost + 1
			for _, dxDy := range [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}} {
				path, exists := partialPathsByPosition[[2][2]int{{x, y},  dxDy}]
				if exists && path.cost < minCost {
					minCost = path.cost
				}
			}
			if minCost < maxCost + 1 {
				fmt.Printf("%-*d", cellWidth, minCost)
			} else {
				fmt.Printf("%-*s", cellWidth, string(lines[y][x]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func computePaths(lines []string, partialPaths []PartialPath) *PartialPath {
	partialPathsByPosition := make(map[[2][2]int]*PartialPath)
	partialPathsByPosition[[2][2]int{partialPaths[0].coord, partialPaths[0].dxDy}] = &partialPaths[0]
	priorityQueue := PriorityQueue{};
	heap.Push(&priorityQueue, &partialPaths[0]);
	heap.Init(&priorityQueue)

	path := heap.Pop(&priorityQueue).(*PartialPath)
	for ; lines[path.coord[1]][path.coord[0]] != 'E'; path = heap.Pop(&priorityQueue).(*PartialPath) {
		x, y := path.coord[0], path.coord[1]
		for _, dxDy := range [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}} {
			dx, dy := dxDy[0], dxDy[1]
			if lines[y + dy][x + dx] != '#' {
				cost := path.cost + 1
				if (dx != path.dxDy[0] || dy != path.dxDy[1]) {
					cost += 1000
				}
				adjacentPath, exists := partialPathsByPosition[[2][2]int{{x+dx, y+dy},{dx,dy}}]
				if exists && adjacentPath.cost > cost {
					adjacentPath.cost = cost
					adjacentPath.predecessors = []*PartialPath{path}
				} else if exists && adjacentPath.cost == cost {
					adjacentPath.predecessors = append(adjacentPath.predecessors, path)
				} else if !exists {
					partialPaths = append(partialPaths, PartialPath{-1, cost, [2]int{x+dx, y+dy}, [2]int{dx, dy}, []*PartialPath{path}})
					adjacentPath := &partialPaths[len(partialPaths)-1]
					partialPathsByPosition[[2][2]int{{x+dx, y+dy},{dx,dy}}] = adjacentPath
					heap.Push(&priorityQueue, adjacentPath)
				}
			}
		}
	}
	return path
}

func star1(lines []string) {
	posX, posY := 1, len(lines) - 2
	partialPaths := []PartialPath{{-1, 0, [2]int{posX, posY}, [2]int{1, 0}, []*PartialPath{}}}

	destinationTile := computePaths(lines, partialPaths)

	// printCostMap(partialPathsByPosition, lines)
	fmt.Println("Shortest path cost: ", destinationTile.cost)
}

func star2(lines []string) {
	posX, posY := 1, len(lines) - 2
	partialPaths := []PartialPath{{-1, 0, [2]int{posX, posY}, [2]int{1, 0}, []*PartialPath{}}}

	destinationTile := computePaths(lines, partialPaths)

	tileHasBeenCounted := map[[2]int]bool{destinationTile.coord: true}
	nextTiles := destinationTile.predecessors
	for len(nextTiles) > 0 {
		nextNextTiles := []*PartialPath{}
		for _, tile := range nextTiles {
			tileHasBeenCounted[tile.coord] = true
			nextNextTiles = append(nextNextTiles, tile.predecessors...)
		}
		nextTiles = nextNextTiles
	}

	// printCostMap(partialPathsByPosition, lines)
	fmt.Println("Number of tiles on shortest paths: ", len(tileHasBeenCounted))
}

func main() {
	// var lines = common.ReadLines("day16.txt")
	var lines = common.ReadLines("day16.txt")
	star1(lines)
	star2(lines)
}