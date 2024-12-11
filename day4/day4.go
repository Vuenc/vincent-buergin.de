package main

import (
	common "aoc2024/aoccommon"
	"fmt"
)

func star1(lines []string) {
	const GOAL = "XMAS"
	var DIRECTIONS = [][]int{{1, 0}, {0, 1}, {1, 1}, {-1, 0}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}}
	var xmasCount = 0
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			for _, direction := range DIRECTIONS {
				for i := 0; i < len(GOAL); i++ {
					var x2 = x + i * direction[0]
					var y2 = y + i * direction[1]
					if (x2 < 0 || x2 >= len(lines[0]) || y2 < 0 || y2 >= len(lines)) {
						break
					}
					if lines[y2][x2] != GOAL[i] {
						break
					}
					if i == len(GOAL) - 1 {
						xmasCount++
					}
				}
			}
		}
	}
	fmt.Println("XMAS count:", xmasCount)
}

func star2(lines []string) {
	const GOAL = "MAS"
	var DIRECTIONS = [][]int{{1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	var mas_centers = common.Array2D(len(lines), len(lines[0]))
	var x_masCount = 0
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			for _, direction := range DIRECTIONS {
				for i := 0; i < len(GOAL); i++ {
					var x2 = x + i * direction[0]
					var y2 = y + i * direction[1]
					if (x2 < 0 || x2 >= len(lines[0]) || y2 < 0 || y2 >= len(lines)) {
						break
					}
					if lines[y2][x2] != GOAL[i] {
						break
					}
					if i == len(GOAL) - 1 {
						mas_centers[x + direction[0]][y + direction[1]]++
						if mas_centers[x + direction[0]][y + direction[1]] == 2 {
							x_masCount++
						}
					}

				}
			}
		}
	}
	fmt.Println("X-MAS count:", x_masCount)
}

func main() {
	var lines = common.ReadLines("day4.txt")
	star1(lines)
	star2(lines)
}