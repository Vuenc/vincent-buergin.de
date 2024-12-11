package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func processFile(filename string) (string) {
	bytes, _ := os.ReadFile(filename)
	
	return string(bytes)
}

func star1(input string) {
    re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	var sum = 0
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		var val1, _ = strconv.Atoi(match[1])
		var val2, _ = strconv.Atoi(match[2])
		sum += val1 * val2
	}
	
    fmt.Println("Sum:", sum)
}

func star2(input string) {
    re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	var sum = 0
	var enabled = true
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		if (match[0] == "do()") {
			enabled = true
		} else if (match[0] == "don't()") {
			enabled = false
		} else if enabled {
			var val1, _ = strconv.Atoi(match[1])
			var val2, _ = strconv.Atoi(match[2])
			sum += val1 * val2
		}
	}
	
    fmt.Println("Sum (with do() and don't()):", sum)
}

func main() {
	var input = processFile("day3.txt")
	star1(input)
	star2(input)
}