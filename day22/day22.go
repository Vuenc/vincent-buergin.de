package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"strconv"
)

func star1(numbers []int) {
	sum := 0
	for _, number := range numbers {
		for range 2000 {
			number = ((number * 64) ^ number) % 16777216
			number = ((number / 32) ^ number) % 16777216
			number = ((number * 2048) ^ number) % 16777216
		}
		sum += number
	}
	fmt.Println("Sum of 2000th numbers:", sum)
}

func star2(numbers []int) {
	sequences := map[[4]int]int{}
	for _, number := range numbers {
		lastChanges := [4]int{-10, -10, -10, -10}
		lastDigit := -1
		currentBuyerSequences := map[[4]int]int{}
		for i := range 2000 {
			digit := number % 10
			number = ((number * 64) ^ number) % 16777216
			number = ((number / 32) ^ number) % 16777216
			number = ((number * 2048) ^ number) % 16777216
			if i > 0 {
				for j := range len(lastChanges) - 1 {
					lastChanges[j] = lastChanges[j+1]
				}
				lastChanges[len(lastChanges)-1] = digit - lastDigit
				// fmt.Println(lastChanges, digit, digit-lastDigit, number)
			}
			if i >= 4 {
				_, exists := currentBuyerSequences[lastChanges]
				if !exists {
					currentBuyerSequences[lastChanges] = digit
				}
			}
			lastDigit = digit
		}
		// fmt.Println()
		for seq, val := range currentBuyerSequences {
			sequences[seq] += val
		}
	}
	numBananas := 0
	for seq, value := range sequences {
		if value > numBananas {
			numBananas = value
			fmt.Println(seq, value)
		}
	}
	fmt.Println("Sum of bananas:", numBananas)
}

func main() {
	var lines = common.ReadLines("day22.txt")
	numbers := []int{}
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	star1(numbers)
	// star1([]int{123})
	star2(numbers)
	// star2([]int{123})
	// star2([]int{1,2,3,2024})
}
