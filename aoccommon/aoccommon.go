package aoccommon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ReadLines(filename string) []string {
	file, _ := os.Open(filename)
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func Array2D(height, width int) [][]int {
	a := make([][]int, height)
	for i := range a {
		a[i] = make([]int, width)
	}
	return a
}

func SplitToInts(input, separator string) []int {
	parts := strings.Split(input, separator)
	var output []int

	for _, val := range(parts) {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Errorf("invalid number: %s", parts[0]))
		}
		output = append(output, num)
	}

	return output
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a % b)
}