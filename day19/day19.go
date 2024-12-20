package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"sort"
	"strings"
)

func patternCanBeCreated(pattern string, pieces []string, memoizationMap *map[int]bool) bool {
	if len(pattern) == 0 {
		return true
	}
	memoizedResult, isMemoized := (*memoizationMap)[len(pattern)]
	if isMemoized {
		return memoizedResult
	}
	for _, piece := range pieces {
		if len(pattern) >= len(piece) && pattern[:len(piece)] == piece {
			if patternCanBeCreated(pattern[len(piece):], pieces, memoizationMap) {
				(*memoizationMap)[len(pattern)] = true
				return true
			}
		}
	}
	(*memoizationMap)[len(pattern)] = false
	return false
}

func star1(patterns []string, pieces []string) {
	sort.Strings(pieces)
	possiblePatterns := 0
	for _, pattern := range patterns {
		memoizationMap := map[int]bool{}
		if patternCanBeCreated(pattern, pieces, &memoizationMap) {
			possiblePatterns++
		}
	}
	fmt.Println("Number of possible patterns:", possiblePatterns)
}

func patternNumberOfOptions(pattern string, pieces []string, memoizationMap *map[int]int) int {
	if len(pattern) == 0 {
		return 1
	}
	memoizedResult, isMemoized := (*memoizationMap)[len(pattern)]
	if isMemoized {
		return memoizedResult
	}
	numberOfOptions := 0
	for _, piece := range pieces {
		if len(pattern) >= len(piece) && pattern[:len(piece)] == piece {
			numberOfOptions += patternNumberOfOptions(pattern[len(piece):], pieces, memoizationMap)
		}
	}
	(*memoizationMap)[len(pattern)] = numberOfOptions
	return numberOfOptions
}

func star2(patterns []string, pieces []string) {
	sort.Strings(pieces)
	numberOfAllPossibleOptions := 0
	for _, pattern := range patterns {
		memoizationMap := map[int]int{}
		numberOfAllPossibleOptions += patternNumberOfOptions(pattern, pieces, &memoizationMap)
	}
	fmt.Println("Number of all possible ways patterns can be created:", numberOfAllPossibleOptions)
}


func main() {
	var lines = common.ReadLines("day19.txt")
	pieces := strings.Split(lines[0], ", ")
	patterns := lines[2:]

	star1(patterns, pieces)
	star2(patterns, pieces)
}