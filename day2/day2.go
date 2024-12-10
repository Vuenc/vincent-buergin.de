package main

import (
	common "aoc2024/aoccommon"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processFile(filename string) ([][]int) {
	file, _ := os.Open(filename)
	
	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var report []int;
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			report = append(report, val)
		}
		reports = append(reports, report)
	}

	return reports
}

func isSafe(report []int) (bool) {
	var ascending = report[1] > report[0];
	for i := 0; i < len(report) - 1; i++ {
		if (report[i+1] > report[i]) != ascending {
			return false
		}
		if common.Abs(report[i+1] - report[i]) > 3 || report[i+1] == report[i] {
			return false;
		}
	}
	return true;
}	
func star1(reports [][]int) {
	var count = 0
	for _, report := range reports {
		if isSafe(report) {
			count++;
		}
	}
	fmt.Println("Safe reports:", count)
}
func star2(reports [][]int) {
	var count = 0
	for _, report := range reports {
		if isSafe(report) {
			count++;
		} else {
			for level := 0; level < len(report); level++ {
				if isSafe(append(append([]int(nil), report[:level]...), report[level+1:]...)) {
					count++;
					break;
				}
			}
		}
	}
	fmt.Println("Safe reports (with dampener):", count)
}

func main() {
	var reports = processFile("day2.txt")

	star1(reports)
	star2(reports)
}