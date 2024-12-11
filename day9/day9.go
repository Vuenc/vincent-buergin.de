package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"strconv"
)

func star1(line string) {
	var numbers []int
	length := 0
	for _, c := range line {
		val, _ := strconv.Atoi(string(c))
		numbers = append(numbers, val)
		length += val
	}
	
	blocks := make([]int, length)
	for i := range blocks {
		blocks[i] = -1
	}
	id := 0
	location := 0
	for i, number := range numbers {
		if (i % 2 == 0) {
			for j := 0; j < number; j++ {
				blocks[location + j] = id
			}
			id += 1
		}
		location += number
	}
	fmt.Println(blocks)

	locationLeft := 0
	locationRight := len(blocks) - 1
	for locationLeft < locationRight {
		for ; blocks[locationLeft] >= 0; locationLeft++ {}
		for ; blocks[locationRight] == -1; locationRight-- {}
		if (locationLeft < locationRight) {
			blocks[locationLeft] = blocks[locationRight];
			blocks[locationRight] = -1;
		}
	}

	checksum := 0
	for i, id := range blocks {
		if (id >= 0) {
			checksum += i * id
		}
	}
	fmt.Println(blocks)
	fmt.Println("Checksum:", checksum)
}

func star2(line string) {
	var numbers []int
	length := 0
	for _, c := range line {
		val, _ := strconv.Atoi(string(c))
		numbers = append(numbers, val)
		length += val
	}
	
	blocks := make([]int, length)
	var gaps [][2]int
	var unmovedFiles [][2]int
	for i := range blocks {
		blocks[i] = -1
	}
	id := 0
	location := 0
	for i, number := range numbers {
		if (i % 2 == 0) {
			for j := 0; j < number; j++ {
				blocks[location + j] = id
			}
			unmovedFiles = append(unmovedFiles, [2]int{location, location+number})
			id += 1
		} else {
			gaps = append(gaps, [2]int{location, location+number})
		}
		location += number
	}

	for fileIndex := len(unmovedFiles) - 1; fileIndex >= 0; fileIndex-- {
		file := unmovedFiles[fileIndex];
		size := file[1] - file[0]
		for gapIndex, gap := range gaps {
			if gap[0] > file[0] {
				break
			}
			if gap[1] - gap[0] >= size {
				for i := 0; i < size; i++ {
					blocks[gap[0] + i] = fileIndex
					blocks[file[0] + i] = -1
				}
				gaps[gapIndex][0] += size
				break
			}
		}
	}

	checksum := 0
	for i, id := range blocks {
		if (id >= 0) {
			checksum += i * id
		}
	}
	fmt.Println(blocks)
	fmt.Println("Checksum (whole files):", checksum)
}

func main() {
	var lines = common.ReadLines("day9.txt")
	line := lines[0]

	star1(line)
	star2(line)
}