package main

import (
	common "aoc2024/aoccommon"
	"fmt"
)

func parseField(lines []string) (int, int, int, int, [][]rune) {
	var guardX, guardY, guardDx, guardDy int
	var data [][]rune
	for y, line := range lines {
		data = append(data, []rune(line))
		for x, char := range data[y] {
			if char == '^' || char == '>' || char == 'v' || char == '<' {
				guardX = x
				guardY = y
				if char == '<' { guardDx = -1 }
				if char == '>' { guardDx = 1 }
				if char == '^' { guardDy = -1 }
				if char == 'V' { guardDy = 1 }
				data[guardY][guardX] = 'x'
			}
		}
	}
	return guardX, guardY, guardDx, guardDy, data
}

func star1(lines []string) {
	guardX, guardY, guardDx, guardDy, data := parseField(lines)
	numDistinctFields := 1
	for guardX + guardDx >= 0 && guardX + guardDx < len(lines[0]) && guardY + guardDy >= 0 && guardY + guardDy < len(lines) {
		if data[guardY + guardDy][guardX + guardDx] != '#' {
			guardX += guardDx
			guardY += guardDy
			if data[guardY][guardX] != 'x' {
				data[guardY][guardX] = 'x'
				numDistinctFields++
			}
		} else {
			guardDx, guardDy = -guardDy, guardDx
		}
	}
	fmt.Println("Number of distinct fields:", numDistinctFields)
}

func copy(original [][]rune) [][]rune {
    copied := make([][]rune, len(original))
    for i := range original {
        // Create a new slice for each inner slice
        copied[i] = append([]rune(nil), original[i]...)
    }
    return copied
}

func setCurrentField(data [][]rune, guardX, guardY, guardDx, guardDy int) {
	if data[guardY][guardX] == '.' {
		if guardDx == 0 {
			data[guardY][guardX] = '|'
		} else {
			data[guardY][guardX] = '-'
		}
	} else {
		if (guardDx == 0 && data[guardY][guardX] == '-') || (guardDy == 0 && data[guardY][guardX] == '|') {
			data[guardY][guardX] = '+'
		}
	}
}

func findLoop(data [][]rune, guardX, guardY, guardDx, guardDy int) bool {
	seenDxDy := make([][][4]bool, len(data));
	for i := range seenDxDy {
		seenDxDy[i] = make([][4]bool, len(data[0]))
	}

	for guardX + guardDx >= 0 && guardX + guardDx < len(data[0]) && guardY + guardDy >= 0 && guardY + guardDy < len(data) {
		if data[guardY + guardDy][guardX + guardDx] != '#' {
			guardX += guardDx
			guardY += guardDy
			setCurrentField(data, guardX, guardY, guardDx, guardDy)
			dxDyIndex := 0
			// (1, 0) -> 3; (-1, 0) -> 2; (0, 1) -> 1; (0, -1) -> 0;
			if guardDx == 0 {
				dxDyIndex += (guardDy + 1) / 2
			} else if guardDy == 0 {
				dxDyIndex += (guardDx + 1) / 2 + 2
			}
			if seenDxDy[guardY][guardX][dxDyIndex] {
				// fmt.Println("Loop:")
				// for _, row := range data {
				// 	fmt.Println(string(row))
				// }
				return true
			}
			seenDxDy[guardY][guardX][dxDyIndex] = true
		} else {
			guardDx, guardDy = -guardDy, guardDx
		}
	}
	// fmt.Println("No Loop:")
	// for _, row := range data {
	// 	fmt.Println(string(row))
	// }
	return false
}

func star2(lines []string) {
	successfulObstaclePositions := make(map[[2]int]struct{})
	guardX, guardY, guardDx, guardDy, data := parseField(lines)
	if guardDx == 0 {
		data[guardY][guardX] = '|'
	} else {
		data[guardY][guardX] = '-'
	}
	numSteps := 0

	for guardX + guardDx >= 0 && guardX + guardDx < len(lines[0]) && guardY + guardDy >= 0 && guardY + guardDy < len(lines) {
		// Simulate movement with an extra obstacle
		if data[guardY + guardDy][guardX + guardDx] == '.' {
			copiedField := copy(data)
			copiedField[guardY + guardDy][guardX + guardDx] = '#'
			if findLoop(copiedField, guardX, guardY, guardDx, guardDy) {
				copiedField = copy(data)
				copiedField[guardY + guardDy][guardX + guardDx] = 'O'
				// fmt.Println("Obstacle:")
				// for _, row := range copiedField {
				// 	fmt.Println(string(row))
				// }
				successfulObstaclePositions[[2]int{guardY + guardDy, guardX + guardDx}] = struct{}{}
			}
			// fmt.Println()
		}

		// Do the normal movement
		if data[guardY + guardDy][guardX + guardDx] != '#' {
			guardX += guardDx
			guardY += guardDy
			setCurrentField(data, guardX, guardY, guardDx, guardDy)
		} else {
			guardDx, guardDy = -guardDy, guardDx
		}
		numSteps ++
	}

	fmt.Println("Number of possible obstacle locations:", len(successfulObstaclePositions))
}

func main() {
	var lines = common.ReadLines("day6.txt")
	star1(lines)
	star2(lines)
}