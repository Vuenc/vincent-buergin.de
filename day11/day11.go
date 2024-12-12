package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"strconv"
)


func star1and2(stones []int, blinkLimit int) {
	numStonesByNumber := make(map[int]int)
	for _, stone := range stones {
		numStonesByNumber[stone]++;
	}

	for i := 0; i < blinkLimit; i++ {
		nextNumStonesByNumber := make(map[int]int)
		nextNumStonesByNumber[1] = numStonesByNumber[0]
		numStonesByNumber[0] = 0
		for stone, count := range numStonesByNumber {
			strRepr := strconv.Itoa(stone)
			if len(strRepr) % 2 == 0 {
				num1, _ := strconv.Atoi(strRepr[:len(strRepr)/2])
				num2, _ := strconv.Atoi(strRepr[len(strRepr)/2:])
				nextNumStonesByNumber[num1] += count
				nextNumStonesByNumber[num2] += count
			} else {
				nextNumStonesByNumber[stone * 2024] += count
			}
		}
		numStonesByNumber = nextNumStonesByNumber
	}

	sum := 0
	for _, count := range numStonesByNumber {
		sum += count
	}
	fmt.Println("Final number of stones:", sum)
}

func star1(stones []int) {
	star1and2(stones, 25)
}

func star2(stones []int) {
	star1and2(stones, 75)
}

func main() {
	var lines = common.ReadLines("day11.txt")
	stones := common.SplitToInts(lines[0], " ")
	star1(stones)
	// star1([]int{125, 17})
	star2(stones)
}