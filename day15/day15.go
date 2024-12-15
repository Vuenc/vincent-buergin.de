package main

import (
	common "aoc2024/aoccommon"
	"fmt"
)

var directionMap = map[rune][2]int{'^': {0, -1}, '<': {-1, 0}, 'v': {0, 1}, '>': {1, 0}}

func printField(field [][]rune) {
	for _, row := range field {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func computeGpsSum(field [][]rune) int {
	gpsSum := 0
	for yField, row := range field {
		for xField, c := range row {
			if c == 'O' || c == '[' {
				gpsSum += xField + 100 * yField
			}
		}
	}
	return gpsSum
}

func findRobot(field [][]rune) (int, int) {
	x := -1
	y := -1
	for yField, row := range field {
		for xField, c := range row {
			if c == '@' {
				x = xField
				y = yField
			}
		}
	}
	return x, y
}

func star1(field [][]rune, instructions []rune) {
	x, y := findRobot(field)
	for _, instruction := range instructions {
		dxDy := directionMap[instruction]
		dx, dy := dxDy[0], dxDy[1]
		var i int
		for i = 1; field[y + dy * i][x + dx * i] == 'O'; i++ {}
		if field[y + dy * i][x + dx * i] == '.' {
			for ; i > 0; i-- {
				field[y + dy * i][x + dx * i] = field[y + dy * (i - 1)][x + dx * (i - 1)]
			}
			field[y][x] = '.'
			y += dy
			x += dx
		}
	}
	fmt.Println("GPS Sum:", computeGpsSum(field))
}

func star2(oldField [][]rune, instructions []rune) {
	var field [][]rune
	// printField(oldField)
	for _, oldRow := range oldField {
		var newRow []rune
		for _, cell := range oldRow {
			if cell != 'O' && cell != '@' {
				newRow = append(newRow, cell, cell)
			} else if cell == 'O' {
				newRow = append(newRow, '[', ']')
			} else {
				newRow = append(newRow, '@', '.')
			}
		}
		field = append(field, newRow)
	}
	// printField(field)

	x, y := findRobot(field)
	for _, instruction := range instructions {
		dxDy := directionMap[instruction]
		dx, dy := dxDy[0], dxDy[1]
		var checkedCoords [][2]int
		coordsToCheck := map[[2]int]bool {{x + dx, y + dy}: true};
		// var emptyCoords [][2]int
		wallFound := false
		for len(coordsToCheck) > 0 && !wallFound {
			nextCoordsToCheck := map[[2]int]bool{}
			for coord := range coordsToCheck {
				if field[coord[1]][coord[0]] == '#' {
					wallFound = true
					break
				} else if field[coord[1]][coord[0]] == '[' {
					nextCoordsToCheck[[2]int{coord[0]+dx, coord[1]+dy}] = true
					if dy != 0 {
						nextCoordsToCheck[[2]int{coord[0]+dx+1, coord[1]+dy}] = true
					}
				} else if field[coord[1]][coord[0]] == ']' {
					nextCoordsToCheck[[2]int{coord[0]+dx, coord[1]+dy}] = true
					if dy != 0 {
						nextCoordsToCheck[[2]int{coord[0]+dx-1, coord[1]+dy}] = true
					}
				} 
				// else if field[coord[1]][coord[0]] == '@' {
				// 	nextCoordsToCheck = append(nextCoordsToCheck, [2]int{coord[0]+dx, coord[1]+dy})
				// }
			}
			for coord := range coordsToCheck {
				checkedCoords = append(checkedCoords, coord)
			}
			coordsToCheck = nextCoordsToCheck
			// fmt.Println("Checked:", checkedCoords)
			// fmt.Println("Next to check:", nextCoordsToCheck)
		}
		if !wallFound {
			for i := len(checkedCoords) - 1; i >= 0; i-- {
				coord := checkedCoords[i]
				field[coord[1]][coord[0]] = field[coord[1] - dy][coord[0] - dx]
				field[coord[1] - dy][coord[0] - dx] = '.'
			}
			y += dy
			x += dx
		}
		// printField(field)
	}
	fmt.Println("GPS Sum (wide field):", computeGpsSum(field))
}

func loadField(lines []string) ([][]rune, []rune) {
	var field [][]rune
	var instructions []rune
	instructionsMode := false
	for _, line := range lines {
		if line == "" {
			instructionsMode = true
		} else if !instructionsMode {
			field = append(field, []rune(line))
		} else {
			instructions = append(instructions, []rune(line)...)
		}
	}
	return field, instructions
}

func main() {
	var lines = common.ReadLines("day15.txt")

	field, instructions := loadField(lines)
	star1(field, instructions)

	field, instructions = loadField(lines)
	star2(field, instructions)
}