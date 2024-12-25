package main

import (
	common "aoc2024/aoccommon"
	"fmt"
)


func star1(keys, locks [][]int) {
	matchingPairs := 0
	for _, key := range keys {
		lockLoop:
		for _, lock := range locks {
			for i := range key {
				if key[i] + lock[i] > 5 {continue lockLoop}
			}
			matchingPairs++
		}
	}
	fmt.Println("Matching key/lock pairs:", matchingPairs)
}

func star2(lines []string) {
}

func main() {
	var lines = common.ReadLines("day25.txt")
	keys := [][]int{}
	locks := [][]int{}
	currentKeyOrLock := []string{}
	for _, line := range append(lines, "") {
		if line != "" {
			currentKeyOrLock = append(currentKeyOrLock, line)
		} else if currentKeyOrLock[0][0] == '#' {
			lock := []int{}
			for x := range len(currentKeyOrLock[0]) {
				for y := len(currentKeyOrLock) - 1; y >= 0; y-- {
					if currentKeyOrLock[y][x] == '#' {
						lock = append(lock, y)
						break
					}
				}
			}
			locks = append(locks, lock)
			currentKeyOrLock = currentKeyOrLock[:0]
		} else {
			key := []int{}
			for x := range len(currentKeyOrLock[0]) {
				for y := range len(currentKeyOrLock) {
					if currentKeyOrLock[y][x] == '#' {
						key = append(key, len(currentKeyOrLock)-1-y)
						break
					}
				}
			}
			keys = append(keys, key)
			currentKeyOrLock = currentKeyOrLock[:0]
		}
	}
	star1(keys, locks)
}