package main

import (
	common "aoc2024/aoccommon"
	"fmt"
)

func star1(antennaLocations map[rune][][2]int, width, height int) {
	antinodesByLocation := make(map[[2]int]int)
	for _, locations := range antennaLocations {
		for _, location1 := range locations {
			for _, location2 := range locations {
				if location1 != location2 {
					antinode := [2]int{location1[0] + 2 * (location2[0] - location1[0]), location1[1] + 2 * (location2[1] - location1[1])}
					if (antinode[0] >= 0 && antinode[1] >= 0 && antinode[0] < width && antinode[1] < height) {
						antinodesByLocation[antinode]++
					}
				}
			}
		}
	}
	fmt.Println("Distinct antinode locations:", len(antinodesByLocation))
}

func star2(antennaLocations map[rune][][2]int, width, height int) {
	antinodesByLocation := make(map[[2]int]int)
	for _, locations := range antennaLocations {
		for i, location1 := range locations {
			for _, location2 := range locations[i+1:] {
				gcd := common.Gcd(common.Abs(location1[0] - location2[0]), common.Abs(location1[1] - location2[1]))
				dx, dy := (location2[0] - location1[0]) / gcd, (location2[1] - location1[1]) / gcd
				x, y := location1[0], location1[1]
				for x >= 0 && y >= 0 && x < width && y < width {
					antinodesByLocation[[2]int{x, y}]++
					x += dx;
					y += dy;
				}
				x, y = location1[0] - dx, location1[1] - dy
				for x >= 0 && y >= 0 && x < width && y < width {
					antinodesByLocation[[2]int{x, y}]++
					x -= dx;
					y -= dy;
				}
			}
		}
	}
	fmt.Println("Distinct antinode locations (taking into account resonant harmonics):", len(antinodesByLocation))
}

func main() {
	var lines = common.ReadLines("day8.txt")
	antennaLocations := make(map[rune][][2]int)
	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				antennaLocations[c] = append(antennaLocations[c], [2]int{x, y})
			}
		}
	}

	star1(antennaLocations, len(lines[0]), len(lines))
	star2(antennaLocations, len(lines[0]), len(lines))
}