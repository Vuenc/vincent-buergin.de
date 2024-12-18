package main

import (
	common "aoc2024/aoccommon"
	"fmt"
)

type FieldState int

const (
	Empty FieldState = iota
	Corrupted
	Used
)

type Cell struct {
	state FieldState
	distanceFromStart int
}

func computePath(field [][]Cell) Cell {	// Don't bother with Dijsktra, let's just do a DFS
	width, height := len(field[0]), len(field)
	startX, startY := 0, 0
	endX, endY := width-1, height-1
	nextCells := [][2]int{{startX, startY}}
	field[startY][startX].state = Used
	field[startY][startX].distanceFromStart = 0
	outerLoop:
	for len(nextCells) > 0 {
		nextNextCells := [][2]int{}
		for _, cell := range nextCells {
			x, y := cell[0], cell[1]
			if x == endX && y == endY {
				break outerLoop
			}
			for _, dxDy := range [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}} {
				dx, dy := dxDy[0], dxDy[1]
				if x + dx >= 0 && x + dx < width && y + dy >= 0 && y + dy < height && field[y+dy][x+dx].state == Empty {
					nextNextCells = append(nextNextCells, [2]int{x+dx, y+dy})
					field[y+dy][x+dx].state = Used
					field[y+dy][x+dx].distanceFromStart = field[y][x].distanceFromStart + 1
				}
			}
		}
		nextCells = nextNextCells
	}
	return field[endY][endX]
}

func printField(field [][]Cell) {
	for _, row := range field {
		for _, c := range row {
			fmt.Print(map[FieldState]string{Empty: ".", Corrupted: "#", Used: "O"}[c.state])
		}
		fmt.Println()
	}
}

func star1(lines []string) {
	field := common.Array2D(71, 71, Cell{Empty, -1})
	for _, line := range lines[:1024] {
		xy := common.SplitToInts(line, ",")
		field[xy[1]][xy[0]].state = Corrupted
	}

	goalCell := computePath(field)
	fmt.Println("Steps needed to reach exit:", goalCell.distanceFromStart)
}

func star2(lines []string) {
	low := 0
	high := len(lines)
	mid := (low + high) / 2
	for low != mid {
		field := common.Array2D(71, 71, Cell{Empty, -1})
		for _, line := range lines[:mid+1] {
			xy := common.SplitToInts(line, ",")
			field[xy[1]][xy[0]].state = Corrupted
		}

		if computePath(field).state == Used {
			low = mid+1
		} else {
			high = mid
		}
		mid = (low + high)/2
	}
	fmt.Println("Coordinates of byte that prevents exit:", lines[mid])
}

func main() {
	var lines = common.ReadLines("day18.txt")
	star1(lines)
	star2(lines)
}
