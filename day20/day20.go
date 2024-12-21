package main

import (
	common "aoc2024/aoccommon"
	"fmt"
)

type FieldState int

const (
	Empty FieldState = iota
	Wall
	EndCell
)

type Cell struct {
	state             FieldState
	distanceFromStart int
}

func computePath(fieldPtr *[][]Cell, startX, startY int) [][2]int { // Don't bother with Dijsktra, let's just do a DFS
	field := *fieldPtr
	width, height := len(field[0]), len(field)
	nextCells := [][2]int{{startX, startY}}
	field[startY][startX].distanceFromStart = 0

	path := [][2]int{}

	for len(nextCells) > 0 {
		nextNextCells := [][2]int{}
		for _, cell := range nextCells {
			x, y := cell[0], cell[1]
			path = append(path, [2]int{x, y})
			if field[y][x].state == EndCell {
				return path
			}
			for _, dxDy := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				dx, dy := dxDy[0], dxDy[1]
				if x+dx >= 0 && x+dx < width && y+dy >= 0 && y+dy < height && field[y+dy][x+dx].state != Wall {
					if field[y+dy][x+dx].distanceFromStart == -1 {
						nextNextCells = append(nextNextCells, [2]int{x + dx, y + dy})
						field[y+dy][x+dx].distanceFromStart = field[y][x].distanceFromStart + 1
					}
				}
			}
		}
		nextCells = nextNextCells
	}
	return path
}

func star1(field [][]Cell, startPos [2]int) {
	path := computePath(&field, startPos[0], startPos[1])
	numberOfCheats := 0
	for _, cell := range path {
		x, y := cell[0], cell[1]
		for _, dxDy := range [][2]int{{-2, 0}, {-1, -1}, {0, -2}, {1, -1}, {2, 0}, {1, 1}, {0, 2}, {-1, 1}} {
			dx, dy := dxDy[0], dxDy[1]
			if x+dx >= 0 && x+dx < len(field[0]) && y+dy >= 0 && y+dy < len(field) {
				if field[y+dy][x+dx].distanceFromStart >= field[y][x].distanceFromStart+102 {
					numberOfCheats++
				}
			}
		}
	}
	fmt.Println("Number of cheats (2 picoseconds rule):", numberOfCheats)
}

func star2(field [][]Cell, startPos [2]int) {
	path := computePath(&field, startPos[0], startPos[1])
	numberOfCheats := 0
	// fmt.Println(path)
	for _, cell := range path {
		x, y := cell[0], cell[1]
		for ddx := 0; ddx <= 20; ddx++ {
			for ddy := 20-ddx; ddy >= 0; ddy-- {
				dxDyRange := [][2]int{{ddx, ddy}, {ddx, -ddy}, {-ddx, ddy}, {-ddx, -ddy}}
				if ddx == 0 && ddy == 0 {
					dxDyRange = [][2]int{{ddx, ddy}}
				} else if ddx == 0 {
					dxDyRange = [][2]int{{ddx, ddy}, {ddx, -ddy}}
				} else if ddy == 0 {
					dxDyRange = [][2]int{{ddx, ddy}, {-ddx, ddy}}					
				}
				for _, dxDy := range dxDyRange {
					dx, dy := dxDy[0], dxDy[1]
					if x+dx >= 0 && x+dx < len(field[0]) && y+dy >= 0 && y+dy < len(field) {
						if field[y+dy][x+dx].distanceFromStart >= field[y][x].distanceFromStart + 100 + common.Abs(dx) + common.Abs(dy) {
							numberOfCheats++
						}
					}
				}

			}
		}
	}
	fmt.Println("Number of cheats (20 picoseconds rule):", numberOfCheats)
}

func loadField(lines []string) ([][]Cell, [2]int) {
	field := common.Array2D(len(lines), len(lines[0]), Cell{})
	startPos := [2]int{-1, -1}
	for y, line := range lines {
		for x, c := range line {
			field[y][x].distanceFromStart = -1
			if c == '.' {
				field[y][x].state = Empty
			} else if c == '#' {
				field[y][x].state = Wall
			} else if c == 'S' {
				startPos = [2]int{x, y}
				field[y][x].distanceFromStart = 0
			} else if c == 'E' {
				field[y][x].state = EndCell
			}
		}
	}
	return field, startPos
}

func main() {
	var lines = common.ReadLines("day20.txt")
	field, startPos := loadField(lines)

	star1(field, startPos)

	field, startPos = loadField(lines)
	star2(field, startPos)
}
