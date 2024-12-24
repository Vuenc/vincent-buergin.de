package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"slices"
	"strconv"
	// "strings"
)

// func pressNumericSequence(sequence string) []rune {
// 	outSequence := []rune{}
// 	positions := map[rune][2]int {
// 		'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
// 		'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
// 		'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
// 		             '0': {1, 3}, 'A': {2, 3},
// 	}
// 	currentPosition := positions['A']
// 	for _, target := range sequence {
// 		targetPosition := positions[target]
// 		for ; currentPosition[0] < targetPosition[0]; currentPosition[0]++ {
// 			outSequence = append(outSequence, '>')
// 		}
// 		for ; currentPosition[1] > targetPosition[1]; currentPosition[1]-- {
// 			outSequence = append(outSequence, '^')
// 		}
// 		for ; currentPosition[1] < targetPosition[1]; currentPosition[1]++ {
// 			outSequence = append(outSequence, 'v')
// 		}
// 		for ; currentPosition[0] > targetPosition[0]; currentPosition[0]-- {
// 			outSequence = append(outSequence, '<')
// 		}
// 		outSequence = append(outSequence, 'A')
// 	}
// 	return outSequence
// }

// func pressArrowSequence(sequence []rune) []rune {
// 	outSequence := []rune{}
// 	positions := map[rune][2]int {
// 		             '^': {1, 0}, 'A': {2, 0},
// 		'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
// 	}
// 	currentPosition := positions['A']
// 	for _, target := range sequence {
// 		targetPosition := positions[target]
// 		if (targetPosition[0] == 0) != (currentPosition[0] == 0) && (targetPosition[1] == 0) != (currentPosition[1] == 0) {
// 			for ; currentPosition[0] < targetPosition[0]; currentPosition[0]++ {
// 				outSequence = append(outSequence, '>')
// 			}
// 			for ; currentPosition[1] < targetPosition[1]; currentPosition[1]++ {
// 				outSequence = append(outSequence, 'v')
// 			}
// 			for ; currentPosition[1] > targetPosition[1]; currentPosition[1]-- {
// 				outSequence = append(outSequence, '^')
// 			}
// 			for ; currentPosition[0] > targetPosition[0]; currentPosition[0]-- {
// 				outSequence = append(outSequence, '<')
// 			}
// 		} else {
// 			for ; currentPosition[0] > targetPosition[0]; currentPosition[0]-- {
// 				outSequence = append(outSequence, '<')
// 			}
// 			for ; currentPosition[1] < targetPosition[1]; currentPosition[1]++ {
// 				outSequence = append(outSequence, 'v')
// 			}
// 			for ; currentPosition[0] < targetPosition[0]; currentPosition[0]++ {
// 				outSequence = append(outSequence, '>')
// 			}
// 			for ; currentPosition[1] > targetPosition[1]; currentPosition[1]-- {
// 				outSequence = append(outSequence, '^')
// 			}
// 		}
// 		outSequence = append(outSequence, 'A')
// 	}
// 	return outSequence
// }

// func star1(lines []string) {
// 	complexitySum := 0
// 	for _, line := range lines {
// 		number, _ := strconv.Atoi(line[:len(line)-1])
// 		arrowSequence1 := pressNumericSequence(line)
// 		arrowSequence2 := pressArrowSequence(arrowSequence1)
// 		arrowSequence3 := pressArrowSequence(arrowSequence2)
// 		// arrowSequence4 := pressArrowSequence(arrowSequence3)

// 		fmt.Println(string(arrowSequence1))
// 		fmt.Println(string(arrowSequence2))
// 		fmt.Println(string(arrowSequence3))
// 		// fmt.Println(string(arrowSequence4))
// 		fmt.Println()

// 		for i, s := range strings.Split(string(arrowSequence3), "A") {
// 			if i >= len(arrowSequence2) {break}
// 			pad := len(s)
// 			if s != "A" {pad++}
// 			fmt.Printf("%-*s ", pad, string(arrowSequence2[i]))
// 		}
// 		fmt.Println()
// 		for _, s := range strings.Split(string(arrowSequence3), "A") {
// 			fmt.Printf(s + "A ")
// 		}
// 		fmt.Println()
// 		fmt.Println(len(arrowSequence3), number)
// 		fmt.Println()
// 		complexitySum += len(arrowSequence3) * number
// 	}
// 	fmt.Println("Sum of complexities:", complexitySum)
// }

var arrowPositions = map[rune][2]int{
	'^': {1, 0}, 'A': {2, 0},
	'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
}

func pressNumericSequence(sequence string, maxLevel int) int {
	numberPositions := map[rune][2]int{
		'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
		'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
		'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
		'0': {1, 3}, 'A': {2, 3},
	}
	currentPositions := [][2]int{numberPositions['A']}
	for range maxLevel {
		currentPositions = append(currentPositions, arrowPositions['A'])
	}
	numPresses := 0
	for _, target := range sequence {
		charNumPresses, _ := pressArrowSequenceQuick(&currentPositions, numberPositions[target], [2]int{0,3}, 0, maxLevel)
		numPresses += charNumPresses
	}
	return numPresses
}

func copyPositions(positions *[][2]int) [][2]int {
	newPositions := make([][2]int, len(*positions))
	for i := 0; i < len(*positions); i++ {
		newPositions[i] = [2]int{(*positions)[i][0], (*positions)[i][1]}
	}
	return newPositions
}

const (
	INVALID rune = '*'
)

func pressArrowSequence(currentPositions *[][2]int, targetPosition [2]int, targetRune rune, forbiddenField [2]int, level int, maxLevel int) []rune {
	if level == maxLevel+1 {
		return []rune{targetRune}
	}

	outSequences := [2][]rune{{}, {}}
	allPositions := [2][][2]int{copyPositions(currentPositions), copyPositions(currentPositions)}

	// fmt.Println("Seq 1")
	positions := allPositions[0]
	for ; positions[level][0] < targetPosition[0]; positions[level][0]++ { // >
		if positions[level] == forbiddenField { outSequences[0] = []rune{INVALID} }
		outSequences[0] = append(outSequences[0], pressArrowSequence(&positions, arrowPositions['>'], '>', [2]int{0,0}, level+1, maxLevel)...)
	}
	for ; positions[level][1] > targetPosition[1]; positions[level][1]-- { // ^
		if positions[level] == forbiddenField { outSequences[0] = []rune{INVALID} }
		outSequences[0] = append(outSequences[0], pressArrowSequence(&positions, arrowPositions['^'], '^', [2]int{0,0}, level+1, maxLevel)...)
	}
	for ; positions[level][1] < targetPosition[1]; positions[level][1]++ { // v
		if positions[level] == forbiddenField { outSequences[0] = []rune{INVALID} }
		outSequences[0] = append(outSequences[0], pressArrowSequence(&positions, arrowPositions['v'], 'v', [2]int{0,0}, level+1, maxLevel)...)
	}
	for ; positions[level][0] > targetPosition[0]; positions[level][0]-- { // <
		if positions[level] == forbiddenField { outSequences[0] = []rune{INVALID} }
		outSequences[0] = append(outSequences[0], pressArrowSequence(&positions, arrowPositions['<'], '<', [2]int{0,0}, level+1, maxLevel)...)
	}
	outSequences[0] = append(outSequences[0], pressArrowSequence(&positions, arrowPositions['A'], 'A', [2]int{0,0}, level+1, maxLevel)...)

	// fmt.Println("Seq 2")
	positions = allPositions[1]
	// secondOptionUsed := false
	// if positions[level][0] > 0 && targetPosition[0] > 0 && positions[level][0] != targetPosition[0] && positions[level][1] != targetPosition[1] {
	if true {
		// secondOptionUsed = true
		for ; positions[level][0] > targetPosition[0]; positions[level][0]-- { // <
			if positions[level] == forbiddenField { outSequences[1] = []rune{INVALID} }
			outSequences[1] = append(outSequences[1], pressArrowSequence(&positions, arrowPositions['<'], '<', [2]int{0,0}, level+1, maxLevel)...)
		}
		for ; positions[level][1] < targetPosition[1]; positions[level][1]++ { // v
			if positions[level] == forbiddenField { outSequences[1] = []rune{INVALID} }
			outSequences[1] = append(outSequences[1], pressArrowSequence(&positions, arrowPositions['v'], 'v', [2]int{0,0}, level+1, maxLevel)...)
		}
		for ; positions[level][1] > targetPosition[1]; positions[level][1]-- { // ^
			if positions[level] == forbiddenField { outSequences[1] = []rune{INVALID} }
			outSequences[1] = append(outSequences[1], pressArrowSequence(&positions, arrowPositions['^'], '^', [2]int{0,0}, level+1, maxLevel)...)
		}
		for ; positions[level][0] < targetPosition[0]; positions[level][0]++ { // >
			if positions[level] == forbiddenField { outSequences[1] = []rune{INVALID} }
			outSequences[1] = append(outSequences[1], pressArrowSequence(&positions, arrowPositions['>'], '>', [2]int{0,0}, level+1, maxLevel)...)
		}
		outSequences[1] = append(outSequences[1], pressArrowSequence(&positions, arrowPositions['A'], 'A', [2]int{0,0}, level+1, maxLevel)...)
	}

	// fmt.Printf("%-*s %d %d\n", 4*level, "", len(outSequences[0]), len(outSequences[1]))
	if len(outSequences[0]) < len(outSequences[1]) || slices.Contains(outSequences[1], INVALID) {
		for i := 0; i < len(*currentPositions); i++ {
			(*currentPositions)[i][0] = allPositions[0][i][0]
			(*currentPositions)[i][1] = allPositions[0][i][1]
		}
		return outSequences[0]
	} else {
		for i := 0; i < len(*currentPositions); i++ {
			(*currentPositions)[i][0] = allPositions[1][i][0]
			(*currentPositions)[i][1] = allPositions[1][i][1]
		}
		return outSequences[1]
	}
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

func trySeq1(currentPositions *[][2]int, targetPosition [2]int, forbiddenField [2]int, level int, maxLevel int) (int, bool) {
	numberPresses := 0
	positions := *currentPositions
	for ; positions[level][0] < targetPosition[0]; positions[level][0]++ { // >
		// fmt.Println("S1 > forbidden field: ", positions[level] == forbiddenField, level, positions[level], targetPosition, forbiddenField)
		if positions[level] == forbiddenField { return -1, false }
		out, ok := pressArrowSequenceQuick(&positions, arrowPositions['>'], [2]int{0,0}, level+1, maxLevel)
		// fmt.Println("S1 > ", level, out, ok)
		if ok { numberPresses += out } else { return -1, false }
	}
	for ; positions[level][1] > targetPosition[1]; positions[level][1]-- { // ^
		// fmt.Println("S1 ^ forbidden field: ", positions[level] == forbiddenField)
		if positions[level] == forbiddenField { return -1, false }
		out, ok := pressArrowSequenceQuick(&positions, arrowPositions['^'], [2]int{0,0}, level+1, maxLevel)
		// fmt.Println("S1 ^ ", level, out, ok)
		if ok { numberPresses += out } else { return -1, false }
	}
	for ; positions[level][1] < targetPosition[1]; positions[level][1]++ { // v
		// fmt.Println("S1 v forbidden field: ", positions[level] == forbiddenField)
		if positions[level] == forbiddenField { return -1, false }
		out, ok := pressArrowSequenceQuick(&positions, arrowPositions['v'], [2]int{0,0}, level+1, maxLevel)
		// fmt.Println("S1 v ", level, out, ok)
		if ok { numberPresses += out } else { return -1, false }
	}
	for ; positions[level][0] > targetPosition[0]; positions[level][0]-- { // <
		// fmt.Println("S1 < forbidden field: ", positions[level] == forbiddenField)
		if positions[level] == forbiddenField { return -1, false }
		out, ok := pressArrowSequenceQuick(&positions, arrowPositions['<'], [2]int{0,0}, level+1, maxLevel)
		// fmt.Println("S1 < ", level, out, ok)
		if ok { numberPresses += out } else { return -1, false }
	}
	out, ok := pressArrowSequenceQuick(&positions, arrowPositions['A'], [2]int{0,0}, level+1, maxLevel)
	// fmt.Println("S1 A ", level, out, ok)
	if ok { numberPresses += out } else { return -1, false }
	return numberPresses, true
}

func trySeq2(currentPositions *[][2]int, targetPosition [2]int, forbiddenField [2]int, level int, maxLevel int) (int, bool) {
	numberPresses := 0
	positions := *currentPositions
	for ; positions[level][0] > targetPosition[0]; positions[level][0]-- { // <
		// fmt.Println("S2 < forbidden field: ", positions[level] == forbiddenField)
		if positions[level] == forbiddenField { return -1, false }
		out, ok := pressArrowSequenceQuick(&positions, arrowPositions['<'], [2]int{0,0}, level+1, maxLevel)
		// fmt.Println("S2 < ", level, out, ok)
		if ok { numberPresses += out } else { return -1, false }
	}
	for ; positions[level][1] < targetPosition[1]; positions[level][1]++ { // v
		// fmt.Println("S2 v forbidden field: ", positions[level] == forbiddenField)
		if positions[level] == forbiddenField { return -1, false }
		out, ok := pressArrowSequenceQuick(&positions, arrowPositions['v'], [2]int{0,0}, level+1, maxLevel)
		// fmt.Println("S2 v ", level, out, ok)
		if ok { numberPresses += out } else { return -1, false }
	}
	for ; positions[level][1] > targetPosition[1]; positions[level][1]-- { // ^
		// fmt.Println("S2 ^ forbidden field: ", positions[level] == forbiddenField)
		if positions[level] == forbiddenField { return -1, false }
		out, ok := pressArrowSequenceQuick(&positions, arrowPositions['^'], [2]int{0,0}, level+1, maxLevel)
		// fmt.Println("S2 ^ ", level, out, ok)
		if ok { numberPresses += out } else { return -1, false }
	}
	for ; positions[level][0] < targetPosition[0]; positions[level][0]++ { // >
		// fmt.Println("S2 > forbidden field: ", positions[level] == forbiddenField)
		if positions[level] == forbiddenField { return -1, false }
		out, ok := pressArrowSequenceQuick(&positions, arrowPositions['>'], [2]int{0,0}, level+1, maxLevel)
		// fmt.Println("S2 > ", level, out, ok)
		if ok { numberPresses += out } else { return -1, false }
	}
	out, ok := pressArrowSequenceQuick(&positions, arrowPositions['A'], [2]int{0,0}, level+1, maxLevel)
	// fmt.Println("S2 A ", level, out, ok)
	if ok { numberPresses += out } else { return -1, false }
	return numberPresses, true
}

func pressArrowSequenceQuick(currentPositions *[][2]int, targetPosition [2]int, forbiddenField [2]int, level int, maxLevel int) (int, bool) {
	if level == maxLevel+1 {
		return 1, true
	}

	positions1, positions2 := copyPositions(currentPositions), copyPositions(currentPositions)
	seq1Presses, seq1Ok := trySeq1(&positions1, targetPosition, forbiddenField, level, maxLevel)
	// seq2Presses, seq2Ok := trySeq2(&positions2, targetPosition, forbiddenField, level, maxLevel)
	seq2Presses, seq2Ok := -1, false

	// fmt.Println(seq1Presses, seq1Ok, seq2Presses, seq2Ok)
	if (seq1Presses < seq2Presses || !seq2Ok) && seq1Ok {
		for i := 0; i < len(*currentPositions); i++ {
			(*currentPositions)[i][0] = positions1[i][0]
			(*currentPositions)[i][1] = positions1[i][1]
		}
		return seq1Presses, true
	} else if seq2Ok {
		for i := 0; i < len(*currentPositions); i++ {
			(*currentPositions)[i][0] = positions2[i][0]
			(*currentPositions)[i][1] = positions2[i][1]
		}
		return seq2Presses, true
	}
	panic(fmt.Sprintf("%d: Neither seq1 nor seq2 ok!", level))
}


func star2(lines []string) {
	complexitySum := 0
	for _, line := range lines {
		number, _ := strconv.Atoi(line[:len(line)-1])
		arrowSequence := pressNumericSequence(line, 15)
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
