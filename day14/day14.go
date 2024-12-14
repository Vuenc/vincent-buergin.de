package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

type robot struct {
	x, y, vx, vy int
}

func star1(robots []robot) {
	WIDTH := 101
	HEIGHT := 103
	var quadrant_sums [2][2]int
	for _, robot := range robots {
		robot.x = common.Mod(robot.x + robot.vx * 100, WIDTH)
		robot.y = common.Mod(robot.y + robot.vy * 100, HEIGHT)
		if robot.x != WIDTH/2 && robot.y != HEIGHT/2 {
			quadrant_sums[robot.y / (HEIGHT / 2 + 1)][robot.x / (WIDTH / 2 + 1)]++
		}
	}
	// fmt.Println(quadrant_sums)
	fmt.Println("Safety factor:", quadrant_sums[0][0] * quadrant_sums[0][1] * quadrant_sums[1][0] * quadrant_sums[1][1])
}

func symmetricRatio(field [][]int) float64 {
	width := len(field[0])
	symmetricCount := 0.
	totalCount := 0.
	for y, row := range field {
		for x := range len(row) {
			if field[y][x] > 0 {
				totalCount += 1;
				if field[y][width - 1 - x] > 0 {
					symmetricCount += 1.
				}
			}
		}
	}
	return symmetricCount / totalCount
}

func solidnessRatio(field [][]int, width, height int) float64 {
	neighborCount := 0.
	totalCount := 0.
	for y, row := range field {
		for x := range len(row) {
			if field[y][x] > 0 {
				for _, dxDy := range [][2]int {{0, -1}, {0, 1}, {-1, 0}, {-1, 1}} {
					dx, dy := dxDy[0], dxDy[1]
					if (x + dx >= 0 && x + dx < width && y + dy >= 0 && y + dy < height) {
						totalCount += 1
						if field[y+dy][x+dx] > 0 {
							neighborCount++;
						}
					}
				}
			}
		}
	}
	return neighborCount / totalCount
}

func showField(field [][]int) {
	for _, row := range field {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func assembleArray(fields [][][]int, width, height, num_columns int) [][]int {
	num_rows := (len(fields) + (num_columns - 1)) / num_columns
	image := common.Array2D(num_rows * height + (num_rows - 1), num_columns * width + (num_columns - 1), 1)
	for i := range len(fields) {
		row := i / num_columns
		column := i % num_columns
		for y := range height {
			for x := range width {
				image[row * (height + 1) + y][column * (width + 1) + x] = fields[i][y][x] 
			}
		}
	}
	return image
}

func saveImage(field [][]int, filename string) {
	img := image.NewGray(image.Rect(0, 0, len(field[0]), len(field)))

	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[0]); x++ {
			img.SetGray(x, y, color.Gray{Y: uint8(field[y][x] * 255)})
		}
	}

	outputFile, _ := os.Create(filename)
	_ = png.Encode(outputFile, img)
}

func star2(robots []robot) {
	WIDTH := 101
	HEIGHT := 103
	// symmetric := false;
	field := common.Array2D[int](HEIGHT, WIDTH)
	var steps int
	// maxSymmetricRatio := 0.
	maxSolidnessRatio := 0.;
	var fields [][][]int
	appended := 0
	for steps = 0; steps < 10000; steps++ {		
		for _, robot := range robots {
			x := common.Mod(robot.x + robot.vx * steps, WIDTH)
			y := common.Mod(robot.y + robot.vy * steps, HEIGHT)
			field[y][x] = 1
		}

		// Failed attempt to find a mirror-symmetric image

		// symmetricRatio := symmetricRatio(field)
		// if symmetricRatio > maxSymmetricRatio {
		// 	maxSymmetricRatio = symmetricRatio
		// 	showField(field)
		// 	fmt.Println("New max symmetric ratio:", maxSymmetricRatio, "at step", steps)
		// }

		// After I found the solution, I added this to find it more systematically
		solidnessRatio := solidnessRatio(field, WIDTH, HEIGHT)
		if solidnessRatio > maxSolidnessRatio {
			maxSolidnessRatio = solidnessRatio
			saveImage(field, fmt.Sprintf(`outputs/step-%d.png`, steps))
			fmt.Println("New max solidness ratio:", maxSolidnessRatio, "at step", steps)
		}

		// One period where obvious structures emerge is HEIGHT, with an offset of 12; the other is WIDTH, with an offset of 35
		if (steps % HEIGHT == 12 || steps % WIDTH == 35) {
			fields = append(fields, field)
			field = common.Array2D[int](HEIGHT, WIDTH)
			if (appended % 20 == 0) {
				fmt.Println()
			}
			fmt.Print(steps+1, " ")
			appended++
		} else {
			for y := range HEIGHT {
				for x := range WIDTH {
					field[y][x] = 0
				}
			}
		}
	}
	saveImage(assembleArray(fields, WIDTH, HEIGHT, 20), "outputs/steps.png")
	// showField(field)
	fmt.Println("Steps:", steps)
	fmt.Println()

}

func main() {
	var lines = common.ReadLines("day14.txt")
    robotRegex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	var robots []robot
	for _, line := range lines {
		new_robot := robot{};
		matched := robotRegex.FindStringSubmatch(line)
		new_robot.x, _ = strconv.Atoi(matched[1]);
		new_robot.y, _ = strconv.Atoi(matched[2]);
		new_robot.vx, _ = strconv.Atoi(matched[3]);
		new_robot.vy, _ = strconv.Atoi(matched[4]);
		robots = append(robots, new_robot)
	}

	star1(robots)
	star2(robots)
}