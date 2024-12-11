package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"strconv"
)

func getHeights(lines []string) [][]int {
	heights := common.Array2D(len(lines) + 2, len(lines[0]) + 2, -1)
	for y, line := range lines {
		y += 1
		for x, c := range line {
			x += 1
			heights[y][x], _ = strconv.Atoi(string(c))
		}
	}
	return heights
}

func star1(heights [][]int) {
	reachableTrailheadsSum := 0
	reachableTrailheads := make([][]common.Set[[2]int], len(heights))
	for i := range reachableTrailheads {
		reachableTrailheads[i] = make([]common.Set[[2]int], len(heights[0]))
	}
	for k := 9; k >= 0; k-- {
		for y := 1; y < len(heights) - 1; y++ {
			row := heights[y]
			for x := 1; x < len(row) - 1; x++ {
				if row[x] == k {
					if k == 9 {
						reachableTrailheads[y][x] = common.NewSet([][2]int{{x, y}})
					} else {
						for _, coords := range [][2]int{{y-1,x}, {y+1,x}, {y, x-1}, {y, x+1}} {
							x2, y2 := coords[1], coords[0]
							if heights[y2][x2] == k + 1 {
								reachableTrailheads[y][x] = common.Union(reachableTrailheads[y][x], reachableTrailheads[y2][x2])
							}
						}
					}
					if k == 0 {
						reachableTrailheadsSum += len(reachableTrailheads[y][x])
					}
				}
			}
		}
	}

	fmt.Println("Sum of reachable trailheads:", reachableTrailheadsSum)
}

func star2(heights [][]int) {
	hikableTrailsCpunt := 0
	numTrailsFromCell := make([][]int, len(heights))
	for i := range numTrailsFromCell {
		numTrailsFromCell[i] = make([]int, len(heights[0]))
	}
	for k := 9; k >= 0; k-- {
		for y := 1; y < len(heights) - 1; y++ {
			row := heights[y]
			for x := 1; x < len(row) - 1; x++ {
				if row[x] == k {
					if k == 9 {
						numTrailsFromCell[y][x] = 1
					} else {
						for _, coords := range [][2]int{{y-1,x}, {y+1,x}, {y, x-1}, {y, x+1}} {
							x2, y2 := coords[1], coords[0]
							if heights[y2][x2] == k + 1 {
								numTrailsFromCell[y][x] += numTrailsFromCell[y2][x2]
							}
						}
					}
					if k == 0 {
						hikableTrailsCpunt += numTrailsFromCell[y][x]
					}
				}
			}
		}
	}

	fmt.Println("Number of hikable trails:", hikableTrailsCpunt)	
}

func main() {
	var lines = common.ReadLines("day10.txt")
	heights := getHeights(lines)
	star1(heights)
	star2(heights)
	// star2(lines)
}