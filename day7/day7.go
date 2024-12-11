package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"strconv"
	"strings"
)

func eval(numbers []int, operators []rune) int {
	sum := 0
	current := numbers[0]
	for i, op := range operators {
		if (op == '*') {
			current *= numbers[i+1]
		} else {
			sum += current
			current = numbers[i+1]
		}
	}
	sum += current
	return sum
}

func evalLeftToRight(numbers []int, operators []rune) int {
	current := numbers[0]
	for i, op := range operators {
		if (op == '*') {
			current *= numbers[i+1]
		} else if (op == '+') {
			current += numbers[i+1]
		} else { // if (op == '|')
			mul := 10
			for mul <= numbers[i+1] {
				mul *= 10
			}
			current = current * mul + numbers[i+1]
		}
	}
	return current
}

func hasAssignmentPlusMul(numbers []int, result int) bool {
	operators := make([]rune, len(numbers) - 1)

	j := 0
	for j >= 0 {
		if operators[j] == 0 {
			operators[j] = '+'
			j++
		} else if operators[j] == '+' {
			operators[j] = '*'
			j++
		} else {
			operators[j] = 0
			j--
		}
		if j == len(operators) {
			if evalLeftToRight(numbers, operators) == result {
				// for k := range operators {
				// 	fmt.Print(numbers[k], string(operators[k]))
				// }
				// fmt.Print(numbers[len(numbers)-1], "=", result)
				// fmt.Println()
				return true
			}
			j--
		}
	}
	return false
}

func hasAssignmentPlusMulConc(numbers []int, result int) bool {
	operators := make([]rune, len(numbers) - 1)

	j := 0
	for j >= 0 {
		if operators[j] == 0 {
			operators[j] = '+'
			j++
		} else if operators[j] == '+' {
			operators[j] = '*'
			j++
		} else if operators[j] == '*' {
			operators[j] = '|'
			j++
		} else {
			operators[j] = 0
			j--
		}
		if j == len(operators) {
			if evalLeftToRight(numbers, operators) == result {
				// for k := range operators {
				// 	fmt.Print(numbers[k], string(operators[k]))
				// }
				// fmt.Print(numbers[len(numbers)-1], "=", result)
				// fmt.Println()
				return true
			}
			j--
		}
	}
	return false
}

func star1(allResults []int, allNumbers [][]int) {
	totalCalibrationResult := 0
	for i := range allResults {
		if hasAssignmentPlusMul(allNumbers[i], allResults[i]) {
			totalCalibrationResult += allResults[i]
		}
	}
	fmt.Println("Total calibration result:", totalCalibrationResult)
}


func star2(allResults []int, allNumbers [][]int) {
	totalCalibrationResult := 0
	for i := range allResults {
		if hasAssignmentPlusMulConc(allNumbers[i], allResults[i]) {
			totalCalibrationResult += allResults[i]
		}
	}
	fmt.Println("Total calibration result (with concatenation):", totalCalibrationResult)
}

func main() {
	var lines = common.ReadLines("day7.txt")

	// fmt.Println(eval([]int{1,2,3,4,5,6,7,8}, []rune{'+', '*', '*', '+', '*', '*', '*'}))
	// fmt.Println(evalLeftToRight([]int{11,6,16,20}, []rune{'+', '*', '+'}))
	// fmt.Println(evalLeftToRight([]int{15, 6}, []rune{'|'}))

	var allResults []int
	var allNumbers [][]int
	for _, line := range lines {
		result, _ := strconv.Atoi(line[:strings.Index(line, ":")])
		var numbers []int
		for _, number := range strings.Split(line[strings.Index(line, ":")+2:], " ") {
			num, _ := strconv.Atoi(number)
			numbers = append(numbers, num)
		}
		allResults = append(allResults, result)
		allNumbers = append(allNumbers, numbers)
	}

	star1(allResults, allNumbers)
	star2(allResults, allNumbers)
}