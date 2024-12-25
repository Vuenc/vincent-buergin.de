package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"strconv"
)

var arrowPositions = map[rune][2]int{
	'X': {0, 0}, '^': {1, 0}, 'A': {2, 0},
	'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
}

var numberPositions = map[rune][2]int{
	'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
	'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
	'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
	'X': {0, 3}, '0': {1, 3}, 'A': {2, 3},
}

func generateLowerLevelSequence(currentPosition [2]int, targetPosition [2]int, forbiddenField [2]int, moveInXDirectionFirst bool) ([]rune, bool) {
	sequence := []rune{}
	if moveInXDirectionFirst {
		for ; currentPosition[0] < targetPosition[0]; currentPosition[0]++ {
			if currentPosition == forbiddenField { return nil, false }
			sequence = append(sequence, '>')
		}
		for ; currentPosition[0] > targetPosition[0]; currentPosition[0]-- {
			if currentPosition == forbiddenField { return nil, false }
			sequence = append(sequence, '<')
		}
	}
	for ; currentPosition[1] < targetPosition[1]; currentPosition[1]++ {
		if currentPosition == forbiddenField { return nil, false }
		sequence = append(sequence, 'v')
	}
	for ; currentPosition[1] > targetPosition[1]; currentPosition[1]-- {
		if currentPosition == forbiddenField { return nil, false }
		sequence = append(sequence, '^')
	}
	if !moveInXDirectionFirst {
		for ; currentPosition[0] < targetPosition[0]; currentPosition[0]++ {
			if currentPosition == forbiddenField { return nil, false }
			sequence = append(sequence, '>')
		}
		for ; currentPosition[0] > targetPosition[0]; currentPosition[0]-- {
			if currentPosition == forbiddenField { return nil, false }
			sequence = append(sequence, '<')
		}
	}
	sequence = append(sequence, 'A')
	return sequence, true
}

func pressSequence(sequence []rune, level int, maxLevel int, positionsRef *map[rune][2]int, lookupRef *[]map[string]int) int {
	if level == maxLevel + 1 {return len(sequence)}
	lookup := (*lookupRef)[level]
	sequenceString := string(sequence)
	lookupValue, exists := lookup[sequenceString]
	if exists {
		return lookupValue
	}

	positions := *positionsRef
	currentPosition := positions['A']
	forbiddenField := positions['X']
	numPresses := 0
	for _, target := range sequence {
		targetPosition := positions[target]
		lowerLevelSequence1, valid1 := generateLowerLevelSequence(currentPosition, targetPosition, forbiddenField, true)
		lowerLevelSequence2, valid2 := generateLowerLevelSequence(currentPosition, targetPosition, forbiddenField, false)
		currentPosition = targetPosition
		numPresses1, numPresses2 := -1, -1
		if valid1 {
			numPresses1 = pressSequence(lowerLevelSequence1, level+1, maxLevel, &arrowPositions, lookupRef)
		}
		if valid2 {
			numPresses2 = pressSequence(lowerLevelSequence2, level+1, maxLevel, &arrowPositions, lookupRef)
		}
		if !valid1 { numPresses += numPresses2 }
		if !valid2 { numPresses += numPresses1 }
		if valid1 && valid2 { numPresses += min(numPresses1, numPresses2) }		
	}
	lookup[sequenceString] = numPresses
	return numPresses
}

func pressNumericSequence(sequence string, maxLevel int) int {
	lookup := make([]map[string]int, maxLevel+1)
	for i := range maxLevel + 1 {
		lookup[i] = map[string]int{}
	}
	return pressSequence([]rune(sequence), 0, maxLevel, &numberPositions, &lookup)
}

func copyPositions(positions *[][2]int) [][2]int {
	newPositions := make([][2]int, len(*positions))
	for i := 0; i < len(*positions); i++ {
		newPositions[i] = [2]int{(*positions)[i][0], (*positions)[i][1]}
	}
	return newPositions
}

func star1(lines []string) {
	complexitySum := 0
	for _, line := range lines {
		number, _ := strconv.Atoi(line[:len(line)-1])
		numPresses := pressNumericSequence(line, 2)
		complexitySum += numPresses * number
	}
	fmt.Println("Sum of complexities:", complexitySum)
}





func star2(lines []string) {
	complexitySum := 0
	for _, line := range lines {
		number, _ := strconv.Atoi(line[:len(line)-1])
		arrowSequence := pressNumericSequence(line, 25)
		complexitySum += arrowSequence * number
	}
	fmt.Println("Sum of complexities:", complexitySum)
}

func main() {
	var lines = common.ReadLines("day21.txt")
	// var lines = common.ReadLines("test.txt")
	star1(lines)
	star2(lines)
}
