package main

import (
	common "aoc2024/aoccommon"
	"fmt"
)


func star1(connectedComponents [][][2]int, componentIdByCell [][]int) {
	perimeterAreaSum := 0
	for componentId, component := range connectedComponents {
		if len(component) == 0 {
			continue
		}
		perimeter := 0
		for _, coord := range component {
			x, y := coord[0], coord[1]
			perimeter += 4
			for _, dxdy := range [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}} {
				dx, dy := dxdy[0], dxdy[1]
				if x + dx >= 0 && x + dx < len(componentIdByCell[0]) && y + dy >= 0 && y + dy < len(componentIdByCell) && componentIdByCell[y+dy][x+dx] == componentId {
					perimeter--;
				}
			}
		}
		perimeterAreaSum += perimeter * len(component)
	}
	fmt.Println("(Perimeter)*(Area) sum:", perimeterAreaSum)
}

func star2(connectedComponents [][][2]int, componentIdByCell [][]int) {
	numSidesAreaSum := 0
	sidesCountedByCell := common.Array2D(len(componentIdByCell), len(componentIdByCell[0]), [4]bool{})
	for componentId, component := range connectedComponents {
		if len(component) == 0 {
			continue
		}
		numSides := 0
		for _, coord := range component {
			x, y := coord[0], coord[1]
			numSides += 4
			for _, dxDySide := range [][3]int{{-1,0, 0}, {1,0, 1}, {0,-1, 2}, {0,1, 3}} {
				dx, dy, side := dxDySide[0], dxDySide[1], dxDySide[2]
				if x + dx >= 0 && x + dx < len(componentIdByCell[0]) && y + dy >= 0 && y + dy < len(componentIdByCell) && componentIdByCell[y+dy][x+dx] == componentId {
					numSides--;
				} else {
					sidesCountedByCell[y][x][side] = true
				}
			}
		}
		for _, coord := range component {
			x, y := coord[0], coord[1]
			for _, dxDySide := range[][3]int{{0, -1, 0}, {0, -1, 1}, {-1, 0, 2}, {-1, 0, 3}} {
				dx, dy, side := dxDySide[0], dxDySide[1], dxDySide[2]
				if x + dx >= 0 && y + dy >= 0 && sidesCountedByCell[y][x][side] && sidesCountedByCell[y+dy][x+dx][side] && componentIdByCell[y+dy][x+dx] == componentId {
					numSides--;
				}
			}
		}
		numSidesAreaSum += numSides * len(component)
	}
	fmt.Println("(Num sides)*(Area) sum:", numSidesAreaSum)
}

func main() {
	var lines = common.ReadLines("day12.txt")
	var connectedComponents [][][2]int
	componentIdByCell := common.Array2D(len(lines), len(lines[0]), -1)
	for y, line := range lines {
		for x := range line {
			eligibleNeighbors := [][2]int{}
			if (x > 0) {
				eligibleNeighbors = append(eligibleNeighbors, [2]int{x-1, y})
			}
			if (y > 0) {
				eligibleNeighbors = append(eligibleNeighbors, [2]int{x, y-1})
			}
			for _, neighbor := range eligibleNeighbors {
				if lines[neighbor[1]][neighbor[0]] == lines[y][x] {
					if componentIdByCell[y][x] == -1 {
						componentIdByCell[y][x] = componentIdByCell[neighbor[1]][neighbor[0]]
						connectedComponents[componentIdByCell[y][x]] = append(connectedComponents[componentIdByCell[y][x]], [2]int{x, y})
					} else if componentIdByCell[y][x] != componentIdByCell[neighbor[1]][neighbor[0]] {
						neighborComponent := componentIdByCell[neighbor[1]][neighbor[0]]
						ownComponent := componentIdByCell[y][x]
						for _, coord := range connectedComponents[componentIdByCell[y][x]] {
							componentIdByCell[coord[1]][coord[0]] = neighborComponent
							connectedComponents[neighborComponent] = append(connectedComponents[neighborComponent], coord)
						}
						connectedComponents[ownComponent] = nil
					}
				}
			}
			// If no eligible neighbors of same char: create new component
			if componentIdByCell[y][x] == -1 {
				componentIdByCell[y][x] = len(connectedComponents)
				connectedComponents = append(connectedComponents, [][2]int{{x, y}})
			}
		}
	}

	star1(connectedComponents, componentIdByCell)
	star2(connectedComponents, componentIdByCell)
}